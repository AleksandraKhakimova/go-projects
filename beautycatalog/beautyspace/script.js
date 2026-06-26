// Данные товаров с указанием изображений
let products = [
    { id: 1, name: "Увлажняющий крем Aravia", category: "Кремы", price: 1500, image: "увлажняющий_крем.webp" },
    { id: 2, name: "Ночной крем CLINIQUE", category: "Кремы", price: 5000, image: "ночной_крем.webp" },
    { id: 3, name: "Дневной крем с SPF ЗINA", category: "Кремы", price: 1800, image: "дневной_крем_спф.jpg" },
    { id: 4, name: "Цветочный парфюм GUCCI", category: "Парфюм", price: 10000, image: "цветочный_парфюм.webp" },
    { id: 5, name: "Древесный парфюм WOOD", category: "Парфюм", price: 4200, image: "древесный_парфюм.webp" },
    { id: 6, name: "Цитрусовый парфюм Guerlain", category: "Парфюм", price: 8700, image: "цитрусовый_парфюм.webp" },
    { id: 7, name: "Матовая помада MAC", category: "Помада", price: 4700, image: "матовая_помада.webp" },
    { id: 8, name: "Глянцевая помада MAYBELLINE", category: "Помада", price: 2900, image: "глянцевая_помада.jpg" },
    { id: 9, name: "Блеск Dior", category: "Помада", price: 5400, image: "жидкая_помада.jpeg" },
    { id: 10, name: "Антивозрастной крем PAYOT", category: "Кремы", price: 2500, image: "антивозрастной_крем.jpg" },
    { id: 11, name: "Фруктовый парфюм TOM FORD", category: "Парфюм", price: 35000, image: "фруктовый_парфюм.jpg" },
    { id: 12, name: "Бальзам для губ EAT MY", category: "Помада", price: 350, image: "бальзам_для_губ.jpg" }
];

let discountActive = false;
let discountPercent = 0;
let cart = [];
let cartDiscountActive = false;

// Вход на сайтfunction enterSite() 
{
    document.getElementById('stars-animation').style.display = 'none';
    document.getElementById('main-site').classList.add('active');
    document.getElementById('main-site').style.display = 'block';
    renderProducts();
    updateStats();
    renderAdminProducts();
    updateCartBadge();
}

// Переключение страниц
function showPage(page) {
    document.querySelectorAll('.page').forEach(p => p.classList.remove('active'));
    document.querySelectorAll('.nav-btn').forEach(b => b.classList.remove('active'));
    
    document.getElementById(`${page}-page`).classList.add('active');
    
    document.querySelectorAll('.nav-btn').forEach(btn => {
        if (btn.textContent.toLowerCase().includes(page === 'catalog' ? 'каталог' : 
            page === 'stats' ? 'статистика' : 'администратор')) {
            btn.classList.add('active');
        }
    });
    
    if (page === 'stats') updateStats();
    if (page === 'admin') renderAdminProducts();
}

// Переключение скидки (вкл/выкл)
function toggleDiscount(checked) {
    discountActive = checked;
    discountPercent = checked ? 10 : 0;
    renderProducts();
    updateStats();
}

// Фильтрация товаров
function filterProducts() {
    const searchTerm = document.getElementById('search-input').value.toLowerCase().trim();
    const sortType = document.getElementById('sort-filter').value;
    const selectedCategories = getSelectedCategories();
    
    let filtered = products.filter(p => {
        const matchesSearch = p.name.toLowerCase().includes(searchTerm);
        const matchesCategory = selectedCategories.length === 0 || selectedCategories.includes(p.category);
        return matchesSearch && matchesCategory;
    });
    
    if (sortType === 'price-asc') {
        filtered.sort((a, b) => a.price - b.price);
    } else if (sortType === 'price-desc') {
        filtered.sort((a, b) => b.price - a.price);
    }
    
    renderProducts(filtered);
}

// Получение выбранных категорий
function getSelectedCategories() {
    const checkboxes = document.querySelectorAll('#category-dropdown input[type="checkbox"]');
    const selected = [];
    checkboxes.forEach(cb => {
        if (cb.checked) selected.push(cb.value);
    });
    return selected;
}

// Переключение дропдауна
function toggleDropdown() {
    const dropdown = document.getElementById('category-dropdown');
    dropdown.classList.toggle('show');
}

document.addEventListener('click', function(e) {
    const dropdown = document.getElementById('category-dropdown');
    const btn = document.querySelector('.dropdown-btn');
    if (!dropdown.contains(e.target) && !btn.contains(e.target)) {
        dropdown.classList.remove('show');
    }
});

// Корзина
function addToCart(productId) {
    const product = products.find(p => p.id === productId);
    if (!product) return;
    
    const existingItem = cart.find(item => item.id === productId);
    if (existingItem) {
        existingItem.quantity += 1;
    } else {
        cart.push({
            id: product.id,
            name: product.name,
            price: product.price,
            image: product.image,
            quantity: 1
        });
    }
    updateCartBadge();
    renderProducts();
    renderCart();
}

function removeFromCart(productId) {
    cart = cart.filter(item => item.id !== productId);
    updateCartBadge();
    renderProducts();
    renderCart();
}

function updateQuantity(productId, delta) {
    const item = cart.find(i => i.id === productId);
    if (!item) return;
    
    item.quantity += delta;
    if (item.quantity <= 0) {
        removeFromCart(productId);
    } else {
        updateCartBadge();
        renderProducts();
        renderCart();
    }
}

function getCartQuantity(productId) {
    const item = cart.find(i => i.id === productId);
    return item ? item.quantity : 0;
}

function updateCartBadge() {
    const total = cart.reduce((sum, item) => sum + item.quantity, 0);
    document.getElementById('cart-count').textContent = total;
}

function toggleCart() {
    const modal = document.getElementById('cart-modal');
    if (modal.classList.contains('show')) {
        modal.classList.remove('show');
    } else {
        renderCart();
        modal.classList.add('show');
    }
}

function closeCart() {
    document.getElementById('cart-modal').classList.remove('show');
}

function renderCart() {
    const container = document.getElementById('cart-items-list');
    const totalElement = document.getElementById('cart-total-price');
    const discountCheckbox = document.getElementById('cart-discount-checkbox');
    
    if (cart.length === 0) {
        container.innerHTML = '<div class="empty-cart">🛒 Корзина пуста</div>';
        totalElement.textContent = '0.00';
        return;
    }
    
    let html = '';
    let total = 0;
    
    cart.forEach(item => {
        const itemTotal = item.price * item.quantity;
        total += itemTotal;
        const discountedPrice = cartDiscountActive ? item.price * 0.9 : item.price;
        
        html += `
            <div class="cart-item">
                <img src="${item.image || 'placeholder.jpg'}" alt="${item.name}" onerror="this.src='placeholder.jpg'">
                <div class="cart-item-info">
                    <div class="name">${item.name}</div>
                    <div class="price">${discountedPrice.toFixed(2)} ₽</div>
                </div>
                <div class="cart-item-controls">
                    <button onclick="updateQuantity(${item.id}, -1)">-</button>
                    <span class="quantity">${item.quantity}</span>
                    <button onclick="updateQuantity(${item.id}, 1)">+</button>
                </div>
                <button class="remove-btn" onclick="removeFromCart(${item.id})">✕</button>
            </div>
        `;
    });
    
    container.innerHTML = html;
    
    // Применяем скидку к итогу если включена
    if (cartDiscountActive) {
        total = total * 0.9;
    }
    totalElement.textContent = total.toFixed(2);
}

function applyCartDiscount(checked) {
    cartDiscountActive = checked;
    renderCart();
}

function checkout() {
    if (cart.length === 0) {
        alert('Корзина пуста!');
        return;
    }
    
    let total = cart.reduce((sum, item) => sum + item.price * item.quantity, 0);
    if (cartDiscountActive) {
        total = total * 0.9;
    }
    
    alert(`✅ Заказ оформлен!\nСумма: ${total.toFixed(2)} ₽\nСпасибо за покупку!`);
    cart = [];
    updateCartBadge();
    renderProducts();
    renderCart();
    closeCart();
}

// Рендер товаров
function renderProducts(productsToRender = null) {
    const grid = document.getElementById('products-grid');
    const items = productsToRender || products;
    
    if (items.length === 0) {
        grid.innerHTML = '<div class="no-products">😔 Товары не найдены</div>';
        return;
    }
    
    grid.innerHTML = items.map(p => {
        let price = p.price;
        let discountLabel = '';
        let displayPrice = price;
        const cartQuantity = getCartQuantity(p.id);
        
        if (discountActive) {
            displayPrice = price * (1 - discountPercent / 100);
            discountLabel = `<span class="product-discount">-${discountPercent}%</span>`;
        }
        
        const imageSrc = p.image ? p.image : 'placeholder.jpg';
        
        let cartControls = '';
        if (cartQuantity > 0) {
            cartControls = `
                <div class="cart-controls">
                    <button onclick="event.stopPropagation(); updateQuantity(${p.id}, -1)">-</button>
                    <span class="quantity">${cartQuantity}</span>
                    <button onclick="event.stopPropagation(); updateQuantity(${p.id}, 1)">+</button>
                </div>
            `;
        } else {
            cartControls = `
                <button class="add-to-cart-btn" onclick="event.stopPropagation(); addToCart(${p.id})">
                    🛒 В корзину
                </button>
            `;
        }
        
        return `
            <div class="product-card">
                ${discountLabel}
                <img src="${imageSrc}" alt="${p.name}" class="product-image" onclick="openProduct(${p.id})" onerror="this.src='placeholder.jpg'">
                <div class="product-name" onclick="openProduct(${p.id})">${p.name}</div>
                <div class="product-category">${p.category}</div>
                <div class="product-price">${displayPrice.toFixed(2)} ₽</div>
                ${cartControls}
            </div>
        `;
    }).join('');
}

// Открытие карточки товара
function openProduct(productId) {
    const p = products.find(prod => prod.id === productId);
    if (!p) return;
    
    const modal = document.getElementById('product-modal');
    let price = p.price;
    if (discountActive) {
        price = p.price * (1 - discountPercent / 100);
    }
    
    document.getElementById('modal-title').textContent = p.name;
    document.getElementById('modal-price').textContent = `💰 Цена: ${price.toFixed(2)} ₽`;
    document.getElementById('modal-category').textContent = `📁 Категория: ${p.category}`;
    
    const modalImage = document.getElementById('modal-image');
    modalImage.src = p.image || 'placeholder.jpg';
    modalImage.onerror = function() {
        this.src = 'placeholder.jpg';
    };
    
    const controlsContainer = document.getElementById('modal-cart-controls');
    const cartQuantity = getCartQuantity(p.id);
    
    if (cartQuantity > 0) {
        controlsContainer.innerHTML = `
            <div class="cart-controls">
                <button onclick="updateQuantity(${p.id}, -1); openProduct(${p.id})">-</button>
                <span class="quantity">${cartQuantity}</span>
                <button onclick="updateQuantity(${p.id}, 1); openProduct(${p.id})">+</button>
            </div>
        `;
    } else {
        controlsContainer.innerHTML = `
            <button class="add-to-cart-btn" onclick="addToCart(${p.id}); openProduct(${p.id})">
                🛒 В корзину
            </button>
        `;
    }
    
    modal.classList.add('show');
}

// Закрытие модального окна
function closeModal() {
    document.getElementById('product-modal').classList.remove('show');
}

// Обновление статистики
function updateStats() {
    const content = document.getElementById('stats-content');
    const total = products.length;
    const categories = {};
    const totalPrice = products.reduce((sum, p) => sum + p.price, 0);
    const avgPrice = total > 0 ? totalPrice / total : 0;
    
    products.forEach(p => {
        categories[p.category] = (categories[p.category] || 0) + 1;
    });
    
    let html = `
        <div class="stats-item">📦 Всего товаров: ${total}</div>
        <div class="stats-item">💰 Средняя цена: ${avgPrice.toFixed(2)} ₽</div>
        <div class="stats-item">🏷️ Всего категорий: ${Object.keys(categories).length}</div>
        <div class="stats-item">📊 Товаров по категориям:</div>
    `;
    
    for (const [category, count] of Object.entries(categories)) {
        html += `<div class="stats-item" style="margin-left: 20px; border-left-color: #C0E7FF;">• ${category}: ${count} шт.</div>`;
    }
    
    if (discountActive) {
        html += `<div class="stats-item" style="border-left-color: #FFBECA;">🎉 Активна скидка ${discountPercent}%</div>`;
    }
    
    content.innerHTML = html;
}

// Добавление товара (админ)
function addProduct() {
    const name = document.getElementById('product-name').value.trim();
    const category = document.getElementById('product-category').value;
    const price = parseFloat(document.getElementById('product-price').value);
    const image = document.getElementById('product-image').value.trim();
    
    if (!name || isNaN(price) || price <= 0) {
        alert('Пожалуйста, заполните все поля корректно!');
        return;
    }
    
    const newProduct = {
        id: products.length > 0 ? Math.max(...products.map(p => p.id)) + 1 : 1,
        name,
        category,
        price,
        image: image || ''
    };
    
    products.push(newProduct);
    renderProducts();
    updateStats();
    renderAdminProducts();
    
    document.getElementById('product-name').value = '';
    document.getElementById('product-price').value = '';
    document.getElementById('product-image').value = '';
    
    alert('✅ Товар успешно добавлен!');
}

// Рендер товаров в админке
function renderAdminProducts() {
    const container = document.getElementById('admin-products');
    if (products.length === 0) {
        container.innerHTML = '<p style="color: #666;">Товаров пока нет</p>';
        return;
    }
    
    container.innerHTML = products.map(p => `
        <div class="admin-product-item">
            <div>
                <strong>${p.name}</strong> - ${p.category} - ${p.price.toFixed(2)} ₽
                ${p.image ? `📷 ${p.image}` : '🖼️ Без фото'}
            </div>
            <button class="delete-btn" onclick="deleteProduct(${p.id})">Удалить</button>
        </div>
    `).join('');
}

// Удаление товара
function deleteProduct(id) {
    if (!confirm('Вы уверены, что хотите удалить этот товар?')) return;
    products = products.filter(p => p.id !== id);
    renderProducts();
    updateStats();
    renderAdminProducts();
}

// Закрытие модальных окон при клике на Escape
document.addEventListener('keydown', function(e) {
    if (e.key === 'Escape') {
        closeModal();
        closeCart();
    }
});

// Инициализация
document.addEventListener('DOMContentLoaded', function() {
    // Автоматическая инициализация
});
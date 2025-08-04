// Простой скрипт для демонстрации
console.log("Скрипт успешно загружен!");

document.addEventListener('DOMContentLoaded', () => {
    const features = document.querySelectorAll('.feature');
    
    features.forEach(feature => {
        feature.addEventListener('click', () => {
            feature.classList.toggle('highlight');
        });
    });
    
    console.log("Добро пожаловать на статический сервер Go!");
});
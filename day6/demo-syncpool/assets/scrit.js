console.log("Sync.Pool demo script loaded!");

// Глобальная функция для запуска тестов
window.runTest = async function(withPool) {
    console.log(`Starting test with pool=${withPool}`);
    
    const button = withPool 
        ? document.querySelector('button:nth-child(2)') 
        : document.querySelector('button:nth-child(1)');
    
    const originalText = button.textContent;
    button.textContent = 'Выполняется...';
    button.disabled = true;
    
    try {
        // Увеличим количество итераций для более заметной разницы
        const iterations = 500000; // Было 10000
        console.log("Sending request to /test?pool=" + withPool + "&iterations=" + iterations);
        
        const response = await fetch(`/test?pool=${withPool}&iterations=${iterations}`);
        
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        
        const result = await response.json();
        console.log("Received result:", result);
        
        updateResults(result);
        updateComparison();
    } catch (error) {
        console.error('Test error:', error);
        alert(`Error: ${error.message}`);
    } finally {
        button.textContent = originalText;
        button.disabled = false;
    }
}

function updateResults(result) {
    console.log(`Updating results for ${result.with_pool ? "with pool" : "without pool"}`);
    
    const containerId = result.with_pool ? 'with-pool' : 'without-pool';
    const container = document.getElementById(containerId);
    
    // Конвертируем наносекунды в миллисекунды
    const timeMs = result.duration.Seconds() * 1000;
    const memoryMB = result.memory / 1024 / 1024;
    
    container.querySelector('.result-content').innerHTML = `
        <strong>Время выполнения:</strong> ${timeMs.toFixed(2)} ms<br>
        <strong>Аллокации:</strong> ${result.allocs.toLocaleString()}<br>
        <strong>Память:</strong> ${memoryMB.toFixed(2)} MB<br>
        <strong>Экономия памяти:</strong> ${result.with_pool ? '✅' : '❌'}
    `;
    
    container.classList.add('active');
    setTimeout(() => container.classList.remove('active'), 2000);
}

function updateComparison() {
    console.log("Updating comparison chart");
    
    const withoutPoolElement = document.querySelector('#without-pool .result-content');
    const withPoolElement = document.querySelector('#with-pool .result-content');
    
    // Если нет данных для сравнения
    if (!withoutPoolElement.textContent.includes('Время') || 
        !withPoolElement.textContent.includes('Время')) {
        console.log("Not enough data for comparison");
        return;
    }
    
    // Извлекаем время из текста
    const withoutTime = parseFloat(withoutPoolElement.textContent.match(/Время выполнения: (\d+\.\d+)/)[1]);
    const withTime = parseFloat(withPoolElement.textContent.match(/Время выполнения: (\d+\.\d+)/)[1]);
    
    const maxTime = Math.max(withoutTime, withTime, 1); // защита от деления на 0
    const withoutWidth = (withoutTime / maxTime) * 100;
    const withWidth = (withTime / maxTime) * 100;
    
    document.querySelector('.without-pool-bar').style.width = `${withoutWidth}%`;
    document.querySelector('.without-pool-bar').textContent = `Без Pool: ${withoutTime.toFixed(2)}ms`;
    
    document.querySelector('.with-pool-bar').style.width = `${withWidth}%`;
    document.querySelector('.with-pool-bar').textContent = `С Pool: ${withTime.toFixed(2)}ms`;
}

// Альтернативная регистрация обработчиков
document.addEventListener('DOMContentLoaded', () => {
    console.log("DOM fully loaded");
    
    document.querySelector('button:nth-child(1)').addEventListener('click', () => runTest(false));
    document.querySelector('button:nth-child(2)').addEventListener('click', () => runTest(true));
});
function runCode() {
    const codeInput = document.getElementById('codeInput').value;
    const outputElement = document.getElementById('output');

    outputElement.textContent = "Выполнение кода...";

    setTimeout(() => {
        if (codeInput.includes('[') && codeInput.includes(']')) {
            outputElement.textContent = "Массив успешно создан!\n[1 2 3 4 5]";
        } else if (codeInput.includes('len')) {
            outputElement.textContent = "Длина массива: 5";
        } else if (codeInput.includes('range')) {
            outputElement.textContent = "Индекс: 0, Значение: 1\nИндекс: 1, Значение: 2\nИндекс: 2, Значение: 3";
        } else {
            outputElement.textContent = "Код выполнен успешно!\nПопробуйте использовать массивы с синтаксисом [5]int{1, 2, 3, 4, 5}";
        }
    }, 1000);
}

document.addEventListener('DOMContentLoaded', function() {
    document.getElementById('codeInput').value =
        '// Пример создания массива\narr := [5]int{1, 2, 3, 4, 5}\nfmt.Println("Массив:", arr)\n\n// Доступ к элементам\nfmt.Println("Первый элемент:", arr[0])\n\n// Изменение элемента\narr[2] = 99\nfmt.Println("После изменения:", arr)';
});

// 1. Конвертация температуры
let celsius = parseFloat(prompt("Введите температуру в Цельсиях:"));
let fahrenheit = (celsius * 9/5) + 32;
console.log(`${celsius}°C = ${fahrenheit}°F`);

// 2. Чётное или нечётное
let num = parseInt(prompt("Введите число:"));
if (num % 2 === 0) {
    console.log("Число чётное");
} else {
    console.log("Число нечётное");
}

// 3. Проверка простого числа
function isPrime(n) {
    if (n <= 1) return false;
    for (let i = 2; i <= Math.sqrt(n); i++) {
        if (n % i === 0) return false;
    }
    return true;
}

let number = parseInt(prompt("Введите число для проверки на простоту:"));
console.log(isPrime(number) ? "Число простое" : "Число не является простым");

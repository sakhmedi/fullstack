# 1. Конвертация температуры
celsius = float(input("Введите температуру в Цельсиях: "))
fahrenheit = (celsius * 9/5) + 32
print(f"{celsius}°C = {fahrenheit}°F")

# 2. Чётное или нечётное
num = int(input("Введите число: "))
if num % 2 == 0:
    print("Число чётное")
else:
    print("Число нечётное")

# 3. Проверка простого числа
num = int(input("Введите число для проверки на простоту: "))
is_prime = True
if num <= 1:
    is_prime = False
else:
    for i in range(2, int(num**0.5)+1):
        if num % i == 0:
            is_prime = False
            break

print("Число простое" if is_prime else "Число не является простым")

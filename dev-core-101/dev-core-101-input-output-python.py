from datetime import datetime

# Ввод данных от пользователя
user_name = input("Enter your name: ")
age = int(input("Enter your age: "))

# Текущий год
current_year = datetime.now().year

# Рассчитываем год, когда пользователю исполнится 100 лет
year_hundred = current_year + (100 - age)

# Выводим результат
print(f"Hello, {user_name}! You will turn 100 years old in the year {year_hundred}.")

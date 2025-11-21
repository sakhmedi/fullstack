# Функция для проверки имени (только буквы)
def get_name():
    while True:
        name = input("Введите ваше имя: ").strip()

        # Проверяем: имя состоит только из букв
        if name.replace(" ", "").isalpha():
            return name
        else:
            print("Имя должно состоять только из букв. Попробуйте снова.\n")

# Приветствие
name = get_name()
print(f"Привет, {name}! Добро пожаловать в мини-викторину!")

# Список вопросов
questions = [
    "Столица Казахстана?",
    "Сколько планет в Солнечной системе?",
    "Сколько континентов на Земле?"
]

# Правильные ответы
answers = [
    ["астана"],
    ["8", "восемь"],
    ["7", "семь"]
]

score = 0  # количество правильных ответов

# Цикл викторины
for i in range(len(questions)):
    print(f"\nВопрос {i+1}: {questions[i]}")
    user_answer = input("Ваш ответ: ").strip().lower()

    if user_answer in answers[i]:
        print("Верно!")
        score += 1
    else:
        print(f"Неверно. Правильный ответ: {answers[i][0]}")

# Итог
print(f"\n{name}, вы ответили правильно на {score} из {len(questions)} вопросов.")

import random

def starting_message():
    print("Welcome to the Number Guessing Game!")
    print("I'm thinking of a number between 1 and 100.")
    print("Please select a difficulty.")
    print("1. Easy (10 chances)")
    print("2. Medium (5 chances)")
    print("3. Hard (3 chances)")
    print("4. Exit")

def get_difficulty():
    while True:
        try:
            difficulty = int(input("Enter your choice: "))
            if difficulty in [1, 2, 3, 4]:
                return difficulty
            else:
                print("Invalid choice. Please select a valid difficulty level.")
        except ValueError:
            print("Invalid input. Please enter a number.")

def play_game_message(difficulty):
    print(f"Great! You have selected the {difficulty} difficulty level.")
    print("Let's play the game!")

def play_game(chances, correct_number):
    for i in range(chances):
        print(f"You have {chances - i} chances left.")
        try:
            user_guess = int(input("Enter your guess: "))
            if user_guess == correct_number:
                print("Congratulations! You have guessed the correct number.")
                break
            elif user_guess < correct_number:
                print("Too low.")
            else:
                print("Too high.")
        except ValueError:
            print("Invalid input. Please enter a number.")
    else:
        print(f"Sorry, you've run out of chances. The correct number was {correct_number}.")

def get_random_number():
    return random.randint(1, 100)

def get_chances(difficulty):
    if difficulty == 1:
        play_game_message("Easy")
        return 10
    elif difficulty == 2:
        play_game_message("Medium")
        return 5
    elif difficulty == 3:
        play_game_message("Hard")
        return 3
    elif difficulty == 4:
        print("Exiting the game...")
        exit()

def main():
    starting_message()
    difficulty = get_difficulty()
    chances = get_chances(difficulty)
    correct_number = get_random_number()
    play_game(chances, correct_number)

if __name__ == "__main__":
    main()


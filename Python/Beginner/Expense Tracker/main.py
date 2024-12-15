import json
from datetime import datetime
import os

def addExpense(data, split, counter):
    description = ''
    amount = ''
    for value in range(0, len(split)-1):
        if split[value] == '--description':
            description = split[value+1]
        if split[value] == '--amount':
            amount = split[value+1]
    expense = {
        "id": counter,
        "description": description,
        "amount": amount,
        "date": datetime.now().strftime("%Y-%m-%d")
    }
    data.append(expense)
    with open('list.json', 'w') as file:
        json.dump(data, file, indent=4)
    print("Expense added successfully (ID: %d)" % counter)

def get_counter(data):
    if len(data) == 0:
        return 1
    else:
        return data[-1]['id'] + 1

def deleteExpense(data, expense_id):
    for expense in data:
        if expense['id'] == expense_id:
            data.remove(expense)
            with open('list.json', 'w') as file:
                json.dump(data, file, indent=4)
            print("Expense deleted successfully (ID: %d)" % expense_id)
            return
    print("Expense not found (ID: %d)" % expense_id)

def expense_summary(data, split):
    total = 0
    month = None
    year = None

    if '--month' in split:
        month = int(split[split.index('--month') + 1])
    if '--year' in split:
        year = int(split[split.index('--year') + 1])

    for expense in data:
        expense_date = datetime.strptime(expense['date'], "%Y-%m-%d")
        if (month is None or expense_date.month == month) and (year is None or expense_date.year == year):
            total += int(expense['amount'])

    if month and year:
        print("Total expenses for month %d of year %d: %d" % (month, year, total))
    elif month:
        print("Total expenses for month %d: %d" % (month, total))
    elif year:
        print("Total expenses for year %d: %d" % (year, total))
    else:
        print("Total expenses: %d" % total)

def main():
    print('Expense Tracker')
    print("""
    CMDs:
    1. To add expense: add --description "Lunch" --amount 20
    2. To show expense list: list
    3. To show summary: summary
    4. To delete expense: delete --id 2
    5. To show summary of specific month: summary --month 8
    6. To exit: exit
    """)

    if os.path.exists('list.json'):
        try:
            with open('list.json', 'r') as file:
                data = json.load(file)
        except json.JSONDecodeError:
            data = []
    else:
        data = []

    while True:
        cmd = input('> ')
        split = cmd.split(' ')

        if split[0] == 'add':
            counter = get_counter(data)
            addExpense(data, split, counter)
        elif split[0] == 'list':
            with open('list.json', 'r') as file:
                data = json.load(file)
            print(json.dumps(data, indent=4))
        elif split[0] == 'summary':
            expense_summary(data, split)
        elif split[0] == 'delete':
            if '--id' in split:
                try:
                    expense_id = int(split[split.index('--id') + 1])
                    deleteExpense(data, expense_id)
                except (ValueError, IndexError):
                    print("Invalid ID")
            else:
                print("Please provide an ID")
        elif split[0] == 'exit':
            break
        else:
            print('Please enter a valid command...')

if __name__ == "__main__":
    main()

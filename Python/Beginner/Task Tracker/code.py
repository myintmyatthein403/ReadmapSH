import json

def create_task(option, data, counter):
    task = " ".join(option[1:])
    task = { "id": counter, "task": task, "status": "not done" }
    data.write(json.dumps(task) + "\n")
    counter += 1
    return data, counter

def show_tasks(data):
    print("Tasks")
    print("-----")
    data.seek(0)  # Move the file pointer to the beginning
    tasks = [json.loads(line) for line in data]
    for task in tasks:
        status = "[x]" if task["status"] == "done" else "[ ]"
        print(f"{status}, Task: {task['task']}, ID: {task['id']}")
    return

def remove_task(option, data):
    data.seek(0)  # Move the file pointer to the beginning
    tasks = [json.loads(line) for line in data]
    for task in tasks:
        if task["id"] == int(option[1]):
            tasks.remove(task)
            break
    with open("tasks.json", "w") as f:
        for task in tasks:
            f.write(json.dumps(task) + "\n")
    return

def done_task(option, data):
    data.seek(0)  # Move the file pointer to the beginning
    tasks = [json.loads(line) for line in data]
    for task in tasks:
        if task["id"] == int(option[1]):
            task["status"] = "done"
            break
    with open("tasks.json", "w") as f:
        for task in tasks:
            f.write(json.dumps(task) + "\n")
    return

def get_counter(data):
    data.seek(0)  # Move the file pointer to the beginning
    tasks = [json.loads(line) for line in data]
    if tasks:
        return max([task["id"] for task in tasks]) + 1
    else:
        return 0

def main():
    print("Welcome to the Task Manager")
    print("\n1. Create task => create <task>")
    print("2. Show tasks => show")
    print("3. Remove task => remove <id>")
    print("4. Done task => done <id>")
    print("5. Exit => exit")

    while True: 
        option = input("Choose an option: ")
        option = option.split(' ')

        if option[0] == "create":
            with open("tasks.json", "a+") as data:
                counter = get_counter(data)
                create_task(option, data, counter)
        elif option[0] == "show":
            with open("tasks.json", "r") as data:
                show_tasks(data)
        elif option[0] == "remove":
            with open("tasks.json", "r+") as data:
                remove_task(option, data)
        elif option[0] == "done":
            with open("tasks.json", "r+") as data:
                done_task(option, data)
        elif option[0] == "exit":
            break
        else:
            print("Invalid option")

if __name__ == "__main__":
    main()

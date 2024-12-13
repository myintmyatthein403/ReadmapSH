import requests


def getActivites(username):
    url = "https://api.github.com/users/" + username + "/events"
    try:
        response = requests.get(url)
        return response.json()
    except:
        return None

def main():
    print("Github Activities")
    print("CMD:\nTo search Activites - github-activity <username>\nTo exit - exit\n") 
    while True:
        option = input(">")
        split = option.split(' ')

        if split[0] == "github-activity":
            print(split)
            if len(split) != 2:
                print('Invalid command\n')
                continue
            
            username = split[1]
            activities = getActivites(username)
            for activity in activities:
                print(f"Type: {activity['type']}, Repository: {activity['repo']}, Date: {activity['created_at']}")

        
        elif split[0] == "exit":
            break

        else:
            print("Invalid command\n")

        
    
if __name__ == "__main__":
    main()

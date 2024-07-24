import random
import datetime

def generate_random_time():
    hour = random.randint(9, 18)
    minute = random.randint(0, 59)
    second = random.randint(0, 59)
    return f"{hour:02d}-{minute:02d}-{second:02d}"

def generate_data(num_records):
    data = []
    for i in range(1, num_records + 1):
        record_id = f"T{i}"
        wd = random.randint(1, 2)
        priority = random.randint(1, 3)
        t1 = generate_random_time()
        t2 = generate_random_time()
        t3 = generate_random_time()
        
        t1_time = datetime.datetime.strptime(t1, "%H-%M-%S")
        t2_time = t1_time + datetime.timedelta(minutes=random.randint(1, 30))
        t3_time = t2_time + datetime.timedelta(minutes=random.randint(1, 30))
        
        t1 = t1_time.strftime("%H-%M-%S")
        t2 = t2_time.strftime("%H-%M-%S")
        t3 = t3_time.strftime("%H-%M-%S")
        
        record = f"id:{record_id},wd:{wd},priority:{priority},t1:{t1},t2:{t2},t3:{t3}"
        data.append(record)
    
    return data

def write_to_file(file_path, data):
    with open(file_path, 'w') as file:
        for record in data:
            file.write(record + '\n')

file_path = 'reports.log'
num_records = 500
data = generate_data(num_records)
write_to_file(file_path, data)

print(f"Generated {num_records} records in {file_path}")

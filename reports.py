import matplotlib.pyplot as plt
from datetime import datetime

def parse_line(line):
    parts = line.strip().split(',')
    data = {}
    for part in parts:
        key, value = part.split(':')
        data[key] = value
    return data

def read_file(file_path):
    all_data = []
    with open(file_path, 'r') as file:
        for line in file:
            parsed_data = parse_line(line)
            all_data.append(parsed_data)
    return all_data

def time_to_seconds(time_str):
    h, m, s = map(int, time_str.split('-'))
    return h * 3600 + m * 60 + s

def calculate_averages(data):
    total_service_time = 0
    total_queue_time = 0
    count = 0
    
    for entry in data:
        t1 = time_to_seconds(entry['t1'])
        t2 = time_to_seconds(entry['t2'])
        t3 = time_to_seconds(entry['t3'])
        
        service_time = t3 - t2
        queue_time = t2 - t1
        
        total_service_time += service_time
        total_queue_time += queue_time
        count += 1
    
    avg_service_time = total_service_time / count
    avg_queue_time = total_queue_time / count
    
    return avg_service_time, avg_queue_time

def plot_arrival_times(data):
    arrival_hours = [int(entry['t1'].split('-')[0]) for entry in data]
    
    hours_count = {hour: 0 for hour in range(24)}
    for hour in arrival_hours:
        hours_count[hour] += 1
    
    hours = list(hours_count.keys())
    counts = list(hours_count.values())
    
    plt.figure(figsize=(10, 6))
    plt.bar(hours, counts, color='skyblue')
    plt.xlabel('Hour of Day')
    plt.ylabel('Number of Arrivals')
    plt.title('Number of Arrivals by Hour of Day')
    plt.xticks(hours)
    plt.grid(axis='y')
    plt.show()

def plot_w_arrivals(data):
    windows = [int(entry['wd']) for entry in data]
    window_count = {window: 0 for window in range(1, 3)} 
    for window in windows:
        window_count[window] += 1
    
    windows = list(window_count.keys())
    print(windows)
    counts = list(window_count.values())
    
    plt.figure(figsize=(10, 6))
    plt.bar(windows, counts, color='lightgreen')
    plt.xlabel('Window Number')
    plt.ylabel('Number of Clients')
    plt.title('Number of Clients by Window')
    plt.xticks(windows)
    plt.grid(axis='y')
    plt.show()


file_path = 'reports.log' 
data = read_file(file_path)
avg_service_time, avg_queue_time = calculate_averages(data)

print(f"Average service time: {avg_service_time} seconds")
print(f"Average queue time: {avg_queue_time} seconds")

plot_arrival_times(data)
plot_w_arrivals(data)

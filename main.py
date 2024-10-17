import cv2

# Define the video file and time ranges (in seconds)
video_path = 'input.mp4'
time_ranges = [(0, 1200), (1200, 2800)]  # Example time ranges

# Open the video file
cap = cv2.VideoCapture(video_path)

# Get video properties
fps = cap.get(cv2.CAP_PROP_FPS)
width = int(cap.get(cv2.CAP_PROP_FRAME_WIDTH))
height = int(cap.get(cv2.CAP_PROP_FRAME_HEIGHT))
fourcc = cv2.VideoWriter_fourcc(*'mp4v')

# Function to write a segment
def write_segment(start_time, end_time, segment_index):
    cap.set(cv2.CAP_PROP_POS_MSEC, start_time)
    out = cv2.VideoWriter(f'py_output_segment_{segment_index}.mp4', fourcc, fps, (width, height))

    while cap.isOpened():
        current_time = cap.get(cv2.CAP_PROP_POS_MSEC)
        ret, frame = cap.read()
        if not ret or current_time > end_time:
            break
        out.write(frame)
    out.release()

# Process each time range
for i, (start, end) in enumerate(time_ranges):
    write_segment(start, end, i)

cap.release()
cv2.destroyAllWindows()

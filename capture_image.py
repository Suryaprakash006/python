import cv2

def capture_image(image_path='static/captured_image.png'):
    # Open the webcam
    cap = cv2.VideoCapture(0)

    # Capture a single frame
    ret, frame = cap.read()

    # Save the captured frame as an image
    cv2.imwrite(image_path, frame)

    # Release the webcam
    cap.release()

    return image_path

if __name__ == "__main__":
    capture_image()

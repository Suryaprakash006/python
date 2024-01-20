from flask import Flask, render_template, request, jsonify, send_file
import os
from PIL import Image
from gtts import gTTS
import pytesseract
from time import sleep
import subprocess

app = Flask(__name__)

os.environ['TESSDATA_PREFIX'] = r'teseract/' 


# Find Tesseract installation path dynamically
def find_tesseract_cmd():
    try:
        result = subprocess.run(['which', 'tesseract'], stdout=subprocess.PIPE, text=True, check=True)
        return result.stdout.strip()
    except subprocess.CalledProcessError:
        return None

tesseract_cmd = find_tesseract_cmd()

if tesseract_cmd:
    pytesseract.pytesseract.tesseract_cmd = tesseract_cmd
else:
    # Set a default path if Tesseract is not found
    pytesseract.pytesseract.tesseract_cmd = r"C:\Program Files\Tesseract-OCR\tesseract.exe"

# Set the TESSDATA_PREFIX dynamically or use the existing one
tessdata_prefix = os.environ.get('TESSDATA_PREFIX', None)
if tessdata_prefix:
    os.environ['TESSDATA_PREFIX'] = tessdata_prefix
else:
    os.environ['TESSDATA_PREFIX'] = r'D:\Downloads'


@app.route('/')
def index():
    return render_template('index.html')

@app.route('/capture', methods=['POST'])
def capture():
    image_path = 'static/captured_image.png'
    os.system('start /min python capture_image.py')  # Assuming you have a separate script for image capture
    sleep(5)  # Wait for the capture to complete
    text = ocr_tamil(image_path)
    return jsonify({'text': text})

@app.route('/text_to_speech', methods=['POST'])
def text_to_speech():
    text = request.form.get('text')
    output_file = 'static/output.mp3'
    tts = gTTS(text=text, lang='ta')
    tts.save(output_file)
    
    # Play the generated audio file using the default system player
    subprocess.run(['start', 'cmd', '/c', f'start {output_file}'], shell=True)
    
    return send_file(output_file, as_attachment=True)

def ocr_tamil(image_path):
    text = pytesseract.image_to_string(Image.open(image_path), lang='tam')
    return text

if __name__ == '__main__':
    app.run(debug=True)

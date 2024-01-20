from flask import Flask, render_template, request, jsonify, send_file
import os
from PIL import Image
from gtts import gTTS
import pytesseract
from time import sleep
import subprocess

app = Flask(__name__)
os.environ['TESSDATA_PREFIX'] = r'teseract/' 
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

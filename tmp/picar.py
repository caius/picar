import RPi.GPIO as GPIO

def setup_drive():
    GPIO.setup(4, GPIO.OUT)
    GPIO.output(4, GPIO.LOW)
    GPIO.setup(17, GPIO.OUT)
    GPIO.output(17, GPIO.LOW)

def forwards():
    GPIO.output(4, GPIO.HIGH)
    GPIO.output(17, GPIO.LOW)

def backwards():
    GPIO.output(4, GPIO.LOW)
    GPIO.output(17, GPIO.HIGH)

def stop():
    GPIO.output(4, GPIO.LOW)
    GPIO.output(17, GPIO.LOW)

def setup_steering():
    GPIO.setup(18, GPIO.OUT)
    GPIO.output(18, GPIO.LOW)
    GPIO.setup(23, GPIO.OUT)
    GPIO.output(23, GPIO.LOW)

def left():
    GPIO.output(18, GPIO.HIGH)
    GPIO.output(23, GPIO.LOW)

def right():
    GPIO.output(18, GPIO.LOW)
    GPIO.output(23, GPIO.HIGH)

def center():
    GPIO.output(18, GPIO.LOW)
    GPIO.output(23, GPIO.LOW)

def setup():
    GPIO.setmode(GPIO.BCM)
    setup_drive()
    setup_steering()

setup()

# Later…

GPIO.cleanup()

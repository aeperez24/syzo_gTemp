int sensorPin = A0;
int sensorValue = 0;  // variable to store the value coming from the sensor

void setup() {
  // put your setup code here, to run once:

Serial.begin(9600);  
}

void loop() {
  // put your main code here, to run repeatedly:
  char readValue[] = "{\"Measurement\":50,\"Name\":\"temp\"}";
  Serial.write(readValue);
  Serial.write("\n");

//  Serial.print("\n");
  delay(100);
  

}

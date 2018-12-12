#include <WiFiClient.h>
#include <ESP8266WiFi.h>
#include <PubSubClient.h>
const char *ssid = "ImanHana";
const char *password = "Tokayah123!";
const char *mqttServer = "192.168.0.117";
const int mqttPort = 1883;
const char *mqttUser = "YourMQTTUsername";
const char *mqttPassword = "YourMQTTPassword";
WiFiClient espClient;
PubSubClient client(espClient);

String content = "";
char character;

void setup()
{
    Serial.begin(115200);
    WiFi.begin(ssid, password);
    while (WiFi.status() != WL_CONNECTED)
    {
        delay(500);
        Serial.println("Connecting to WiFi..");
    }
    Serial.println("Connected to the WiFi network");
    client.setServer(mqttServer, mqttPort);
    client.setCallback(callback);
    while (!client.connected())
    {
        Serial.println("Connecting to MQTT...");
        if (client.connect("ESP8266Client", mqttUser, mqttPassword))
        {
            Serial.println("connected");
        }
        else
        {
            Serial.println("failed state ");
            Serial.println(client.state());
            delay(2000);
        }
    }
    client.publish("kucing", "Hello ESP World");
    client.subscribe("kucing"); // here is where you later add a wildcard
}

void callback(char *topic, byte *payload, unsigned int length)
{
    Serial.println("Messageved in topic: ");
    Serial.println(topic);
    Serial.println("Message");
    for (int i = 0; i < length; i++)
    {
        character = (char)payload[i];
        content.concat(character);
        //    Serial.print((char)payload[i]);
    }
    if (content != "")
    {
        Serial.println(content);
    }
    content = "";
}

void loop()
{
    client.loop();
}
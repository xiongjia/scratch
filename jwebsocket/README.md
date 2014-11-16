# jwebsocket

A sample Echo Server & Client for [Java WebSocket](https://github.com/TooTallNate/Java-WebSocket)

## NOTES
* The code of this example plugin is based on the Example code in JMeter source.
* This plugin is built and tested on JMeter v2.11.
* JMeter v2.11 is using Java 1.6 and that is the reason of I set the Java v1.6 Compatibility in build.gradle.

## Dependencies
* Java Development Environment
* Gradle
* Maven 
* Eclipse (Option)

## Building
* Builds the server & client: `gradle build` 
* (Option) Imports the project to Eclipse : `gradle eclipse` & Import the folder to your Eclipse

### Build the fat JAR (One .jar file with all dependencies)
* Run Task: `gradle fatJar` (The target .jar file is jwebsocket-{version}-all.jar. )
* Check the usage of this jar file: `java -jar jwebsocket-{version}-all.jar -help`

## Usage
1. Start the Echo Server: `gradle startServ`
2. Launch the Echo client: `gradle sendMsg`
3. Check the output logs of Server & Client

## Note
* The default Echo Server working port is 9005.


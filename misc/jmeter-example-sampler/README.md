# jmeter-example-sampler

An Example plugin for [Apache JMeter](http://jmeter.apache.org/) .

## NOTES
* The code of this example plugin is based on the Example code in JMeter source.
* This plugin is built and tested on JMeter v2.11.
* JMeter v2.11 is using Java 1.6 and that is the reason of I set the Java v1.6 Compatibility in build.gradle.

## Dependencies
* Java Development Environment
* Gradle
* Maven 
* Eclipse (Option)

## Building and Installing
* Builds the plugin .jar file : `gradle jar` 
* Adds the plugin to your JMeter :    
  We just need to copy the "build\libs\jmeter-example-sampler-0.1.0.jar" to your "JMeter\lib\ext" folder. 
* (Option) Imports the project to Eclipse : `gradle eclipse` & Import the folder to your Eclipse

## Usage
1. Creates your Test Plan and Thread Group. 
2. Adds the "My Example" Sampler to your Thread group. 
3. Updates "Response data" in the Sampler configuration panel.
4. Runs this Test Plan.
5. Checks the results report and your should find the "Response data" in the Sampler result.


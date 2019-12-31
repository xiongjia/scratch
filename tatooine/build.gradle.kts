plugins {
  java
}

java {
  sourceCompatibility = JavaVersion.VERSION_1_8
  targetCompatibility = JavaVersion.VERSION_1_8
}

allprojects {
  repositories {
    mavenLocal()
    mavenCentral()
  }
}

subprojects {
  buildDir = file("${rootProject.projectDir}/build/${project.name}")
}


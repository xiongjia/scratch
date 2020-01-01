plugins {
  java
}

java {
  sourceCompatibility = JavaVersion.VERSION_1_8
  targetCompatibility = JavaVersion.VERSION_1_8
}

allprojects {
  apply(plugin = "idea")

  repositories {
    mavenLocal()
    mavenCentral()
  }
}

subprojects {
  buildDir = file("${rootProject.projectDir}/build/${project.name}")
  project.version = "1.0"
}

dependencies {
  testCompile("org.junit.jupiter:junit-jupiter-api:5.5.2")
}


allprojects {
  repositories {
    mavenLocal()
    mavenCentral()
  }
}


subprojects {
  buildDir = file("${rootProject.projectDir}/build/${project.name}")
}


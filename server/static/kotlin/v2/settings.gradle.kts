rootProject.name = "v2"
include("hello")
include("user")
include("protobuf")

pluginManagement {
    repositories {
        mavenLocal()
        gradlePluginPortal()
        google()
        mavenCentral()
    }
    resolutionStrategy {
        eachPlugin {
            val plugin = requested.id.id
            val module = when {
                plugin.startsWith("com.squareup.wire") -> "com.squareup.wire:wire-gradle-plugin:${requested.version}"
                else -> return@eachPlugin
            }
            println("resolutionStrategy for plugin=$plugin : $module")
            useModule(module)
        }
    }
}
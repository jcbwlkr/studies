function generate() {
  mvn archetype:generate -DgroupId=com.jcbwlkr.studies -DartifactId=${1} -DarchetypeArtifactId=maven-archetype-quickstart -DarchetypeVersion=1.4 -DinteractiveMode=false
  cd $1
  sed -i '' -e 's/1.7/1.8/g'
}

function build() {
  mvn package
}

function run() {
  java -cp target/jsonClient-1.0-SNAPSHOT.jar com.jcbwlkr.studies.App
}

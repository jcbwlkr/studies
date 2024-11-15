function generate() {
  mvn archetype:generate -DgroupId=com.jcbwlkr.studies -DartifactId=${1} -DarchetypeArtifactId=maven-archetype-quickstart -DarchetypeVersion=1.4 -DinteractiveMode=false
  cd $1
  sed -i -e 's/1.7/1.8/g' pom.xml
}

function build() {
  mvn package
}

function run() {
  mvn -q exec:java -Dexec.mainClass=com.jcbwlkr.studies.App
}

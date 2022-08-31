/*
Разработайте пакет для чтения конфигурации типичного веб-приложения через флаги или переменные окружения.
Пример конфигурации можно посмотреть здесь. По желанию вы можете задать другие имена полям, сгруппировать их или добавить собственные поля.
Помимо чтения конфигурации приложение также должно валидировать её - например, все URL’ы должны соответствовать ожидаемым форматам.
Работу с конфигурацией необходимо вынести в отдельный пакет (не в пакет main).

*/

/*
Пример конфигурации 
port: 8080
db_url: postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable
jaeger_url: http://jaeger:16686
sentry_url: http://sentry:9000
kafka_broker: kafka:9092
some_app_id: testid
some_app_key: testkey
*/


package main
import (
"flag"
"log"
"os"

"github.com/kelseyhightower/envconfig"
"config"
)
func main() {
	var c config.SrvConfigs
	c.AddSrvConfigs()
	config.SrvConfigs(&c)

}

package config

const SMSFile string = `simulator\sms.data`         //Адрес файла для загрузки данных SMS
const BillingFile string = `simulator\billing.data` //Адрес файла для загрузки данных Billing
const EmailFile string = `simulator\email.data`     //Адрес файла для загрузки данных Email
const VoiceFile string = `simulator\voice.data`     //Адрес файла для загрузки данных Voice

const MMSWeb string = `http://localhost:8383/mms`            //Адрес загрузки данных MMS
const SupportWeb string = `http://localhost:8383/support`    //Адрес загрузки данных Support
const IncidentWeb string = `http://localhost:8383/accendent` //Адрес загрузки данных Incident

const ServerWeb string = `localhost:8282` //Адрес локального сервера для запуска

const TimeTicket float32 = 60.0 / 18.0 //Среднее время обработки тикета

const TimeOut int64 = 30 //Максимальный таймаут в секундах для получения данных

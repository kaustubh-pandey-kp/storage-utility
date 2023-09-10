package constants

type ArtifactType int

const (
	LogArtifact ArtifactType = iota
	ScreenshotArtifact
	VideoArtifact
	WASABI_FAILURE_COUNT_REDIS_KEY = "WASABI_FAILURE_COUNT"
)

type WasabiConfigParams struct {
	WasabiBucketRegion      string
	WasabiAccessKey         string
	WasabiSecretKey         string
	WasabiBucketName        string
	WasabiNetworkBucketName string
	SumoLogicMagicString    string
	WasabiEndpoint          string
}

var (
	WASABI_ENDPOINT_US             = ""
	WASABI_ENDPOINT_EU             = ""
	WASABI_BUCKET                  = ""
	WASABI_NETWORK_BUCKET          = ""
	SUMOLOGIC_MAGIC_STRING         = ""
	SCREENSHOT_ARTEFACT_EXTENSTION string
	VIDEO_ARTEFACT_EXTENSION       string
	LOG_ARTEFACT_EXTENSION         string
	FILE_UPLOAD_MAX_ATTEMPTS       = 10
	FILE_UPLOAD_MAX_TIME_MILLISEC  = 60000
	REDIS_CONNECTION_POOL          = 40
	CACHED_REDIS_HOST              = ""
)

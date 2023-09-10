package cachedRedis

import "github.com/kaustubh-pandey-kp/storage-utility/constants"

func IncrementWasabiFailureCounter() error {
	conn := Pool.Get()
	defer conn.Close()

	// Increment the value in Redis using INCRBY
	_, err := conn.Do("INCRBY", constants.WASABI_FAILURE_COUNT_REDIS_KEY, 1)
	if err != nil {
		return err
	}

	return nil
}
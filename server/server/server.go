package server

func Init() {
	r := Router()
	r.Run("3000")
}

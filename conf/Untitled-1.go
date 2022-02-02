/* ---
 * Add flight records in bulk
 * --- */
func (c APIV1) NewFlightBulk() revel.Result {
	if c.Request.Method == "POST" {
		logfile, err := utils.SaveLogFile(c.Params.Files["logfile"])
		if err != nil {
			// TODO: Log the error
			return ServerError("A server error occurred")
		}
		fmt.Println(logfile)

		// csvFile, _ := os.Open(logfile)
		// r := csv.NewReader(csvFile)
		// for {
		// 	record, err := r.Read()
		// 	if err == io.EOF {
		// 		break
		// 	}
		// 	if err != nil {
		// 		log.Fatal(err)
		// 	}
		// 	fmt.Println(record)
		// }
		return nil
	}
	return MethodNotAllowed("")
}

POST /api/v1/flights 								APIV1.NewFlightBulk
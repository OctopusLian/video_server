/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-05-05 22:09:05
 * @LastEditors: neozhang
 * @LastEditTime: 2022-05-05 22:10:48
 */
package main

import (
	"net/http"
	"video_server/scheduler/dbops"

	"github.com/julienschmidt/httprouter"
)

func vidDelRecHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid-id")

	if len(vid) == 0 {
		sendResponse(w, 400, "video id should not be empty")
		return
	}

	err := dbops.AddVideoDeletionRecord(vid)
	if err != nil {
		sendResponse(w, 500, "Internal server error")
		return
	}

	sendResponse(w, 200, "")
	return
}

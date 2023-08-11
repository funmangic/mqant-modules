// Copyright 2014 loolgame Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package room

import "time"

/*
*
table超时处理机制
*/
type TimeOutTable struct {
	subtable              SubTable
	timeout               int64 //默认超时时间单位秒
	lastCommunicationDate int64
}

func (this *TimeOutTable) TimeOutTableInit(subtable SubTable, timeout int64) {
	this.subtable = subtable
	this.timeout = timeout
	this.lastCommunicationDate = time.Now().Unix()
}
func (this *TimeOutTable) ResetTimeOut() {
	this.lastCommunicationDate = time.Now().Unix()
}

/*
*
检查整个table是否已超时

检查规则:
1. 所有玩家离开场景
2. 场景中没有人超过指定时间
*/
func (this *TimeOutTable) CheckTimeOut() {
	if len(this.subtable.GetSeats()) > 0 {
		this.ResetTimeOut()
		return
	}
	//for _, player := range this.subtable.GetSeats() {
	//	if player != nil {
	//		if this.lastCommunicationDate < player.GetLastReqResDate() {
	//			this.lastCommunicationDate = player.GetLastReqResDate()
	//		}
	//	}
	//}
	if this.timeout > 0 {
		if time.Now().Unix() > (this.lastCommunicationDate + this.timeout) {
			this.subtable.OnTimeOut()
		}
	}
}

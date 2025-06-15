/*
 * Copyright (c) 2025-present the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ipaddrs

import (
	"net"
	"os"
)

const LocalIpEnvKey = "LOCAL_IP"

var localIp = ""

// ----------------------------------------------------------------

func init() {
	localIp, _ = parseLocalIP()
}

// ----------------------------------------------------------------

func LocalIp() string {
	return localIp
}

// ----------------------------------------------------------------

func parseLocalIP() (string, error) {
	if ip, ok := os.LookupEnv(LocalIpEnvKey); ok {
		return ip, nil
	}

	ip := ""
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

LABEL:
	for _, it := range interfaces {
		if ((it.Flags & net.FlagUp) != 0) && ((it.Flags & net.FlagLoopback) == 0) {
			addr, err := it.Addrs()
			if err != nil {
				return "", err
			}
			for _, address := range addr {
				if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
					ip = ipNet.IP.String()
					break LABEL
				}
			}
		}
	}

	return ip, nil
}

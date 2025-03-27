package main

import (
	"log"
)

func updateIpgroup() {
	iKuai, err := loginToIkuai()
	if err != nil {
		log.Println("登录爱快失败：", err)
		return
	}

	for _, ipGroup := range conf.IpGroup {

		preIds, err := iKuai.GetIpGroup(ipGroup.Name)
		if err != nil {
			log.Println("ip分组== 获取准备更新的IP分组列表失败：", ipGroup.Name, err)
			continue
		} else {
			log.Println("ip分组== 获取准备更新的IP分组列表成功", ipGroup.Name, preIds)
		}

		err = updateIpGroup(iKuai, ipGroup.Name, ipGroup.URL, preIds)
		if err != nil {
			log.Printf("ip分组== 添加IP分组'%s@%s'失败：%s\n", ipGroup.Name, ipGroup.URL, err)
		} else {
			log.Printf("ip分组== 添加IP分组'%s@%s'成功\n", ipGroup.Name, ipGroup.URL)
		}

	}

	for _, streamIpPort := range conf.StreamIpPort {
		preIds, err := iKuai.GetStreamIpPortIds(streamIpPort.IpGroup)
		if err != nil {
			log.Println("端口分流== 获取准备更新的端口分流列表失败：", streamIpPort.IpGroup, err)
			continue
		} else {
			log.Println("端口分流== 获取准备更新的端口分流列表成功", streamIpPort.IpGroup, preIds)
		}

		err = updateStreamIpPort(
			iKuai,
			preIds,
			streamIpPort.Type,
			streamIpPort.Interface,
			streamIpPort.Nexthop,
			streamIpPort.SrcAddr,
			streamIpPort.IpGroup,
			streamIpPort.Mode,
			streamIpPort.IfaceBand,
		)
		if err != nil {
			log.Printf("端口分流== 添加端口分流 '%s@%s' 失败：%s\n",
				streamIpPort.Interface+streamIpPort.Nexthop,
				streamIpPort.IpGroup,
				err,
			)
		} else {
			log.Printf("端口分流== 添加端口分流 '%s@%s' 成功\n",
				streamIpPort.Interface+streamIpPort.Nexthop,
				streamIpPort.IpGroup,
			)
		}
	}
}

func updateIpv6group() {
	iKuai, err := loginToIkuai()
	if err != nil {
		log.Println("登录爱快失败：", err)
		return
	}
	
	for _, ipv6Group := range conf.Ipv6Group {
		preIds, err := iKuai.GetIpv6Group(ipv6Group.Name)
		if err != nil {
			log.Println("ipv6分组== 获取准备更新的IPv6分组列表失败：", ipv6Group.Name, err)
			continue
		} else {
			log.Println("ipv6分组== 获取准备更新的IPv6分组列表成功", ipv6Group.Name, preIds)
		}

		err = updateIpv6Group(iKuai, ipv6Group.Name, ipv6Group.URL, preIds)
		if err != nil {
			log.Printf("ipv6分组== 添加IPV6分组'%s@%s'失败：%s\n", ipv6Group.Name, ipv6Group.URL, err)
		} else {
			log.Printf("ipv6分组== 添加IPV6分组'%s@%s'成功\n", ipv6Group.Name, ipv6Group.URL)
		}
	}
}

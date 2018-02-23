package main

import (
    "github.com/zhuyie/goenet"
    "log"
)

type Player struct {
    ID          int64
    PacketCount int64
}

func main() {
    if goenet.Initialize() == 0 {
        defer goenet.Deinitialize()

        address := &goenet.ENetAddress{}
        address.SetHost(goenet.ENET_HOST_ANY)
        address.SetPort(5555)

        event := &goenet.ENetEvent{}

        server := goenet.NewHost(address, 32, 2, 0, 0)
        if server == nil {
            panic("Server Initialization Error")
        } else {
            defer server.Destroy()
        }

        log.Print("Server started\n")

        var nextPlayerID int64
        allPlayers := make(map[int64]*Player)

        for {
            for server.Service(event, 1000) > 0 {
                switch event.EventType() {
                case goenet.ENET_EVENT_TYPE_CONNECT:
                    peer := event.Peer()

                    player := &Player{}
                    player.ID = nextPlayerID
                    nextPlayerID++
                    allPlayers[player.ID] = player

                    peer.SetData(uint(player.ID))
                    log.Printf("Client connected: %d\n", peer.ConnectID())
                    break

                case goenet.ENET_EVENT_TYPE_RECEIVE:
                    peer := event.Peer()

                    playerID := int64(peer.Data())
                    player := allPlayers[playerID]
                    player.PacketCount++

                    length := event.Packet().DataLength()
                    packetData := string(event.Packet().Data())
                    channel := event.ChannelID()
                    log.Printf("packet - length: %d, data: %s, channel: %d", length, packetData, channel)
                    peer.Send(channel, goenet.NewPacket([]byte(packetData), length, goenet.ENET_PACKET_FLAG_RELIABLE))
                    event.Packet().Destroy() // clean up
                    break

                case goenet.ENET_EVENT_TYPE_DISCONNECT:
                    peer := event.Peer()

                    playerID := int64(peer.Data())
                    player := allPlayers[playerID]
                    delete(allPlayers, playerID)

                    log.Printf("Client disconnected: %d PacketCount=%d\n", peer.ConnectID(), player.PacketCount)
                    break
                }
            }
        }
    }
}

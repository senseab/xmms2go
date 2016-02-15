package xmms2go

/*
We need import some enums from C.
Go should not use contents of native C directly
*/

/* IPC object type */
const (
	IPCObjectSignal = iota
	IPCObjectMain
	IPCObjectPlaylist
	IPCObjectConfig
	IPCObjectPlayback
	IPCObjectMedialib
	IPCObjectCollection
	IPCObjectVisualization
	IPCObjectMediaInfoReader
	IPCObjectXform
	IPCObjectBindata
	IPCObjectCollSync
	IPCObjectCourier
	IPCObjectIPCManager
	IPCObjectEnd
)

/* IPC signals */
const (
	IPCSignalPlaylistChanged = iota
	IPCSignalConfigValueChanged
	IPCSignalPlaybackStatus
	IPCSignalPlaybackVolumeChanged
	IPCSignalPlaybackPlaytime
	IPCSignalPlaybackCurrentID
	IPCSignalPlaylistCurrentPos
	IPCSignalPlaylistLoaded
	IPCSignalMedialibEntryAdded
	IPCSignalMedialibEntryUpdate
	IPCSignalMedialibEntryRemoved
	IPCSignalCollectionChanged
	IPCSignalQuit
	IPCSignalMediaInfoReaderStstus
	IPCSignalMediaInfoReaderUnindexed
	IPCSignalCourierMessage
	IPCSignalIPCManagerClientConnected
	IPCSignalIPCManagerClientDisconnected
	IPCSignalEnd
)

const ipcCmdFirst = 32

/* IPC command reply */
const (
	IPCCmdReply = iota
	IPCCmdError
)

/* Signal subsystem methods */
const (
	IPCCmdSignal = ipcCmdFirst + iota
	IPCCmdBroadcast
)

/* Main methods */
const (
	IPCCmdHello = ipcCmdFirst + iota
	IPCCmdQuit
	IPCCmdPluginList
	IPCCmdStats
)

/* Playlist methods */
const (
	IPCCmdReplace = ipcCmdFirst + iota
	IPCCmdSetpOS
	IPCCmdSetpOSRel
	IPCCmdAddURL
	IPCCmdAddColl
	IPCCmdRemoveEntry
	IPCCmdMoveEntry
	IPCCmdList
	IPCCmdCurrentPos
	IPCCmdCurrentActive
	IPCCmdInsertURL
	IPCCmdInsertColl
	IPCCmdLoad
	IPCCmdRAdd
	IPCCmdRInsert
)

/* Config methods */
const (
	IPCCmdGetValue = ipcCmdFirst + iota
	IPCCmdSetValue
	IPCCmdRegValue
	IPCCmdListValues
)

/* Playback methods */
const (
	IPCCmdStart = ipcCmdFirst + iota
	IPCCmdStop
	IPCCmdPausE
	IPCCmdDecoderKill
	IPCCmdCPlaytime
	IPCCmdSeekMS
	IPCCmdSeekSamples
	IPCCmdPlaybackStatus
	IPCCmdCurrentID
	IPCCmdVolumeSet
	IPCCmdVolumeGet
)

/* Medialib methods */
const (
	IPCCmdInfo = ipcCmdFirst + iota
	IPCCmdPathImport
	IPCCmdRehash
	IPCCmdGetID
	IPCCmdRemoveID
	IPCCmdPropertySetStr
	IPCCmdPropertySetInt
	IPCCmdPropertyRemove
	IPCCmdMoveURL
	IPCCmdMlibAddURL
)

/* Coll sync methods */
const (
	IPCCmdCollSyncSync = ipcCmdFirst + iota
)

/* Collection methods */
const (
	IPCCmdCollectionGet = ipcCmdFirst + iota
	IPCCmdCollectionList
	IPCCmdCollectionSave
	IPCCmdCollectionRemove
	IPCCmdCollectionFind
	IPCCmdCollectionRename
	IPCCmdQuery
	IPCCmdQueryInfos
	IPCCmdIDListFromPLS
)

/* bindata methods */
const (
	IPCCmdGetData = ipcCmdFirst + iota
	IPCCmdAddData
	IPCCmdRemoveData
	IPCCmdListData
)

/* visualization methods */
const (
	IPCCmdVisualizationQueryVersion = ipcCmdFirst + iota
	IPCCmdVisualizationRegister
	IPCCmdVisualizationInitSHM
	IPCCmdVisualizationInitUDP
	IPCCmdVisualizationProperty
	IPCCmdVisualizationProperties
	IPCCmdVisualizationShutdown
)

/* xform methods */
const (
	IPCCmdBrowse = ipcCmdFirst + iota
)

/* courier methods */
const (
	IPCCmdSendMessage = ipcCmdFirst + iota
	IPCCmdReplyMessage
	IPCCmdGetConnectedClients
)

/* Playlist changed events */
const (
	PlaylistChangedAdd = iota
	PlaylistChangedInsert
	PlaylistChangedShuffle /* deprecated */
	PlaylistChangedRemove
	PlaylistChangedClear /* deprecated */
	PlaylistChangedMove
	PlaylistChangedSort /* deprecated */
	PlaylistChangedUpdate
	PlaylistChangedReplace
)

/* Collection changed events */
const (
	CollectionChangedAdd = iota
	CollectionChangedUpdate
	CollectionChangedRename
	CollectionChangedRemove
)

/* Playlist current ID methods */
const (
	PlaylistCurrentIDForget = iota
	PlaylistCurrentIDKeep
	PlaylistCurrentIDMoveToFront
)

/* Playback status */
const (
	PlaybackStatusStop = iota
	PlaybackStatusPlay
	PlaybackStatusPause
)

/* Playback seek */
const (
	PlaybackSeekCur = 1 + iota
	PlaybackSeekSet
)

/* Media info reader status */
const (
	MediaInfoReaderStatusIdle = iota
	MediaInfoReaderStatusRunning
)

/* Plugin type */
const (
	PluginTypeAll = iota
	PluginTypeOutput
	PluginTypeXform
)

/* Collection type */
const (
	CollectionTypeReference = iota
	CollectionTypeUniverse
	CollectionTypeUnion
	CollectionTypeIntersection
	CollectionTypeComplement
	CollectionTypeHas
	CollectionTypeMatch
	CollectionTypeToken
	CollectionTypeEquals
	CollectionTypeNotEqual
	CollectionTypeSmaller
	CollectionTypeSmallerEQ
	CollectionTypeGreater
	CollectionTypeGreaterEQ
	CollectionTypeOrder
	CollectionTypeLimit
	CollectionTypeMediaset
	CollectionTypeIDList
	CollectionTypeLast = CollectionTypeIDList
)

/* Medialib entry status */
const (
	MedialibEntryStatusNew = iota
	MedialibEntryStatusOK
	MedialibEntryStatusResolving
	MedialibEntryStatusNotAvailable
	MedialibEntryStatusRehash
)

/* Log level */
const (
	LogLevelUnknown = iota
	LogLevelFatal
	LogLevelFail
	LogLevelError
	LogLevelInfo
	LogLevelDebug
	LogLevelCount /* must be last */
)

/* C2C reply policy */
const (
	C2CReplyPolicyNoReply = iota
	C2CReplyPolicySingleReply
	C2CReplyPolicyMultiReply
)

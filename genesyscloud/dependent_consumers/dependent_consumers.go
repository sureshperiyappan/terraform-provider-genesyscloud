package dependent_consumers

var dependentConsumerMap map[string]string
var flowTypeObjectMap map[string]string

func SetDependentObjectMaps() map[string]string {
	if len(dependentConsumerMap) < 1 {
		dependentConsumerMap = make(map[string]string)
		dependentConsumerMap["QUEUE"] = "genesyscloud_routing_queue"
		dependentConsumerMap["INBOUNDCALLFLOW"] = "genesyscloud_flow"
		dependentConsumerMap["USER"] = "genesyscloud_user"
		dependentConsumerMap["GROUP"] = "genesyscloud_group"
		dependentConsumerMap["OUTBOUNDCALLFLOW"] = "genesyscloud_flow"
		dependentConsumerMap["KNOWLEDGEBASEDOCUMENT"] = "genesyscloud_knowledge_document"
		dependentConsumerMap["LANGUAGE"] = "genesyscloud_routing_language"
		dependentConsumerMap["ACDSKILL"] = "genesyscloud_routing_skill"
		dependentConsumerMap["ACDWRAPUPCODE"] = "genesyscloud_routing_wrapupcode"
		dependentConsumerMap["CONTACTLIST"] = "genesyscloud_outbound_contactlist"
		dependentConsumerMap["DATATABLE"] = "genesyscloud_architect_datatable"
		dependentConsumerMap["FLOWOUTCOME"] = "genesyscloud_flow_outcome"
		dependentConsumerMap["EMAILROUTE"] = "genesyscloud_routing_email_route"
		dependentConsumerMap["EMERGENCYGROUP"] = "genesyscloud_architect_emergencygroup"
		dependentConsumerMap["GRAMMAR"] = "genesyscloud_architect_grammar"
		dependentConsumerMap["INBOUNDCHATFLOW"] = "genesyscloud_flow"
		dependentConsumerMap["INBOUNDEMAILFLOW"] = "genesyscloud_flow"
		dependentConsumerMap["INBOUNDSHORTMESSAGEFLOW"] = "genesyscloud_flow"
		dependentConsumerMap["INQUEUEEMAILFLOW"] = "genesyscloud_flow"
		dependentConsumerMap["INQUEUESHORTMESSAGEFLOW"] = "genesyscloud_flow"
		dependentConsumerMap["FLOWMILESTONE"] = "genesyscloud_flow_milestone"
		dependentConsumerMap["IVRCONFIGURATION"] = "genesyscloud_architect_ivr"
		dependentConsumerMap["KNOWLEDGEBASE"] = "genesyscloud_knowledge_knowledgebase"
		dependentConsumerMap["OAUTHCLIENT"] = "genesyscloud_oauth_client"
		dependentConsumerMap["RECORDINGPOLICY"] = "genesyscloud_recording_media_retention_policy"
		dependentConsumerMap["RESPONSE"] = "genesyscloud_responsemanagement_response"
		dependentConsumerMap["SCHEDULE"] = "genesyscloud_architect_schedules"
		dependentConsumerMap["SCHEDULEGROUP"] = "genesyscloud_architect_schedulegroups"
		dependentConsumerMap["USERPROMPT"] = "genesyscloud_architect_user_prompt"
		dependentConsumerMap["WIDGET"] = "genesyscloud_widget_deployment"
		dependentConsumerMap["COMPOSERSCRIPT"] = "genesyscloud_script"
		dependentConsumerMap["ACDLANGUAGE"] = "genesyscloud_routing_language"
		dependentConsumerMap["INQUEUECALLFLOW"] = "genesyscloud_flow"
		dependentConsumerMap["DATAACTION"] = "genesyscloud_integration_action"

		dependentConsumerMap["SECURECALLFLOW"] = "genesyscloud_flow"
		dependentConsumerMap["VOICEFLOW"] = "genesyscloud_flow"
		dependentConsumerMap["WORKFLOW"] = "genesyscloud_flow"
		dependentConsumerMap["WORKITEMFLOW"] = "genesyscloud_flow"
		dependentConsumerMap["BOTFLOW"] = "genesyscloud_flow"
		dependentConsumerMap["COMMONMODULEFLOW"] = "genesyscloud_flow"
		dependentConsumerMap["VOICEMAILFLOW"] = "genesyscloud_flow"
		dependentConsumerMap["SURVEYINVITEFLOW"] = "genesyscloud_flow"
	}

	return dependentConsumerMap
}

func SetFlowTypeObjectMaps() map[string]string {
	if len(flowTypeObjectMap) < 1 {
		flowTypeObjectMap = make(map[string]string)
		flowTypeObjectMap["BOT"] = "BOTFLOW"
		flowTypeObjectMap["COMMONMODULE"] = "COMMONMODULEFLOW"
		flowTypeObjectMap["DIGITALBOT"] = "DIGITALBOTFLOW"
		flowTypeObjectMap["INBOUNDCALL"] = "INBOUNDCALLFLOW"
		flowTypeObjectMap["INBOUNDCHAT"] = "INBOUNDCHATFLOW"
		flowTypeObjectMap["INBOUNDEMAIL"] = "INBOUNDEMAILFLOW"
		flowTypeObjectMap["INBOUNDSHORTMESSAGE"] = "INBOUNDSHORTMESSAGEFLOW"
		flowTypeObjectMap["INQUEUECALL"] = "INQUEUECALLFLOW"
		flowTypeObjectMap["INQUEUEEMAIL"] = "INBOUNDEMAILFLOW"
		flowTypeObjectMap["INQUEUESHORTMESSAGE"] = "INQUEUESHORTMESSAGEFLOW"
		flowTypeObjectMap["OUTBOUNDCALL"] = "OUTBOUNDCALLFLOW"
		flowTypeObjectMap["SECURECALL"] = "SECURECALLFLOW"
		flowTypeObjectMap["SURVEYINVITE"] = "SURVEYINVITEFLOW"
		flowTypeObjectMap["VOICE"] = "VOICEFLOW"
		flowTypeObjectMap["VOICEMAIL"] = "VOICEMAILFLOW"
		flowTypeObjectMap["WORKFLOW"] = "WORKFLOW"
		flowTypeObjectMap["WORKITEM"] = "WORKITEMFLOW"
	}
	return flowTypeObjectMap
}

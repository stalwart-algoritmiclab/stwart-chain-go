import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgRevoke } from "./types/cosmos/authz/v1beta1/tx";
import { GenericAuthorization } from "./types/cosmos/authz/v1beta1/authz";
import { GrantAuthorization } from "./types/cosmos/authz/v1beta1/authz";
import { GenesisState } from "./types/cosmos/authz/v1beta1/genesis";
import { QueryGrantsResponse } from "./types/cosmos/authz/v1beta1/query";
import { QueryGranterGrantsResponse } from "./types/cosmos/authz/v1beta1/query";
import { MsgExec } from "./types/cosmos/authz/v1beta1/tx";
import { MsgExecResponse } from "./types/cosmos/authz/v1beta1/tx";
import { Grant } from "./types/cosmos/authz/v1beta1/authz";
import { GrantQueueItem } from "./types/cosmos/authz/v1beta1/authz";
import { EventRevoke } from "./types/cosmos/authz/v1beta1/event";
import { QueryGranteeGrantsResponse } from "./types/cosmos/authz/v1beta1/query";
import { MsgGrantResponse } from "./types/cosmos/authz/v1beta1/tx";
import { MsgRevokeResponse } from "./types/cosmos/authz/v1beta1/tx";
import { QueryGrantsRequest } from "./types/cosmos/authz/v1beta1/query";
import { QueryGranteeGrantsRequest } from "./types/cosmos/authz/v1beta1/query";
import { MsgGrant } from "./types/cosmos/authz/v1beta1/tx";
import { EventGrant } from "./types/cosmos/authz/v1beta1/event";
import { QueryGranterGrantsRequest } from "./types/cosmos/authz/v1beta1/query";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/cosmos.authz.v1beta1.MsgRevoke", MsgRevoke],
    ["/cosmos.authz.v1beta1.GenericAuthorization", GenericAuthorization],
    ["/cosmos.authz.v1beta1.GrantAuthorization", GrantAuthorization],
    ["/cosmos.authz.v1beta1.GenesisState", GenesisState],
    ["/cosmos.authz.v1beta1.QueryGrantsResponse", QueryGrantsResponse],
    ["/cosmos.authz.v1beta1.QueryGranterGrantsResponse", QueryGranterGrantsResponse],
    ["/cosmos.authz.v1beta1.MsgExec", MsgExec],
    ["/cosmos.authz.v1beta1.MsgExecResponse", MsgExecResponse],
    ["/cosmos.authz.v1beta1.Grant", Grant],
    ["/cosmos.authz.v1beta1.GrantQueueItem", GrantQueueItem],
    ["/cosmos.authz.v1beta1.EventRevoke", EventRevoke],
    ["/cosmos.authz.v1beta1.QueryGranteeGrantsResponse", QueryGranteeGrantsResponse],
    ["/cosmos.authz.v1beta1.MsgGrantResponse", MsgGrantResponse],
    ["/cosmos.authz.v1beta1.MsgRevokeResponse", MsgRevokeResponse],
    ["/cosmos.authz.v1beta1.QueryGrantsRequest", QueryGrantsRequest],
    ["/cosmos.authz.v1beta1.QueryGranteeGrantsRequest", QueryGranteeGrantsRequest],
    ["/cosmos.authz.v1beta1.MsgGrant", MsgGrant],
    ["/cosmos.authz.v1beta1.EventGrant", EventGrant],
    ["/cosmos.authz.v1beta1.QueryGranterGrantsRequest", QueryGranterGrantsRequest],
    
];

export { msgTypes }
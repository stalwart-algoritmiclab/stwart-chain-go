import { GeneratedType } from "@cosmjs/proto-signing";
import { QueryParamsResponse } from "./types/stwartchain/core/query";
import { MsgRefReward } from "./types/stwartchain/core/tx";
import { QueryAllStatsRequest } from "./types/stwartchain/core/query";
import { DailyStats } from "./types/stwartchain/core/daily_stats";
import { MsgFeesResponse } from "./types/stwartchain/core/tx";
import { QueryGetStatsRequest } from "./types/stwartchain/core/query";
import { MsgSend } from "./types/stwartchain/core/tx";
import { QueryParamsRequest } from "./types/stwartchain/core/query";
import { MsgUpdateParams } from "./types/stwartchain/core/tx";
import { GenesisState } from "./types/stwartchain/core/genesis";
import { MsgUpdateParamsResponse } from "./types/stwartchain/core/tx";
import { MsgIssueResponse } from "./types/stwartchain/core/tx";
import { QueryGetStatsResponse } from "./types/stwartchain/core/query";
import { QueryAllStatsResponse } from "./types/stwartchain/core/query";
import { MsgFees } from "./types/stwartchain/core/tx";
import { MsgIssue } from "./types/stwartchain/core/tx";
import { Stats } from "./types/stwartchain/core/stats";
import { MsgWithdraw } from "./types/stwartchain/core/tx";
import { MsgRefund } from "./types/stwartchain/core/tx";
import { MsgRefRewardResponse } from "./types/stwartchain/core/tx";
import { Params } from "./types/stwartchain/core/params";
import { QueryGetStatsByDateRequest } from "./types/stwartchain/core/query";
import { QueryModulesAddressesRequest } from "./types/stwartchain/core/query";
import { MsgWithdrawResponse } from "./types/stwartchain/core/tx";
import { MsgSendResponse } from "./types/stwartchain/core/tx";
import { MsgRefundResponse } from "./types/stwartchain/core/tx";
import { MsgBurn } from "./types/stwartchain/core/tx";
import { QueryModulesAddressesResponse } from "./types/stwartchain/core/query";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/stwartchain.core.QueryParamsResponse", QueryParamsResponse],
    ["/stwartchain.core.MsgRefReward", MsgRefReward],
    ["/stwartchain.core.QueryAllStatsRequest", QueryAllStatsRequest],
    ["/stwartchain.core.DailyStats", DailyStats],
    ["/stwartchain.core.MsgFeesResponse", MsgFeesResponse],
    ["/stwartchain.core.QueryGetStatsRequest", QueryGetStatsRequest],
    ["/stwartchain.core.MsgSend", MsgSend],
    ["/stwartchain.core.QueryParamsRequest", QueryParamsRequest],
    ["/stwartchain.core.MsgUpdateParams", MsgUpdateParams],
    ["/stwartchain.core.GenesisState", GenesisState],
    ["/stwartchain.core.MsgUpdateParamsResponse", MsgUpdateParamsResponse],
    ["/stwartchain.core.MsgIssueResponse", MsgIssueResponse],
    ["/stwartchain.core.QueryGetStatsResponse", QueryGetStatsResponse],
    ["/stwartchain.core.QueryAllStatsResponse", QueryAllStatsResponse],
    ["/stwartchain.core.MsgFees", MsgFees],
    ["/stwartchain.core.MsgIssue", MsgIssue],
    ["/stwartchain.core.Stats", Stats],
    ["/stwartchain.core.MsgWithdraw", MsgWithdraw],
    ["/stwartchain.core.MsgRefund", MsgRefund],
    ["/stwartchain.core.MsgRefRewardResponse", MsgRefRewardResponse],
    ["/stwartchain.core.Params", Params],
    ["/stwartchain.core.QueryGetStatsByDateRequest", QueryGetStatsByDateRequest],
    ["/stwartchain.core.QueryModulesAddressesRequest", QueryModulesAddressesRequest],
    ["/stwartchain.core.MsgWithdrawResponse", MsgWithdrawResponse],
    ["/stwartchain.core.MsgSendResponse", MsgSendResponse],
    ["/stwartchain.core.MsgRefundResponse", MsgRefundResponse],
    ["/stwartchain.core.MsgBurn", MsgBurn],
    ["/stwartchain.core.QueryModulesAddressesResponse", QueryModulesAddressesResponse],
    
];

export { msgTypes }
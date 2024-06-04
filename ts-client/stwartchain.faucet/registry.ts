import { GeneratedType } from "@cosmjs/proto-signing";
import { QueryParamsRequest } from "./types/stwartchain/faucet/query";
import { QueryGetTokensRequest } from "./types/stwartchain/faucet/query";
import { QueryGetTokensResponse } from "./types/stwartchain/faucet/query";
import { GenesisState } from "./types/stwartchain/faucet/genesis";
import { MsgUpdateTokensResponse } from "./types/stwartchain/faucet/tx";
import { MsgDeleteTokensResponse } from "./types/stwartchain/faucet/tx";
import { QueryAllTokensRequest } from "./types/stwartchain/faucet/query";
import { QueryAllTokensResponse } from "./types/stwartchain/faucet/query";
import { MsgUpdateTokens } from "./types/stwartchain/faucet/tx";
import { MsgDeleteTokens } from "./types/stwartchain/faucet/tx";
import { MsgIssueResponse } from "./types/stwartchain/faucet/tx";
import { QueryParamsResponse } from "./types/stwartchain/faucet/query";
import { MsgUpdateParams } from "./types/stwartchain/faucet/tx";
import { MsgUpdateParamsResponse } from "./types/stwartchain/faucet/tx";
import { MsgCreateTokensResponse } from "./types/stwartchain/faucet/tx";
import { Tokens } from "./types/stwartchain/faucet/tokens";
import { Params } from "./types/stwartchain/faucet/params";
import { MsgIssue } from "./types/stwartchain/faucet/tx";
import { MsgCreateTokens } from "./types/stwartchain/faucet/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/stwartchain.faucet.QueryParamsRequest", QueryParamsRequest],
    ["/stwartchain.faucet.QueryGetTokensRequest", QueryGetTokensRequest],
    ["/stwartchain.faucet.QueryGetTokensResponse", QueryGetTokensResponse],
    ["/stwartchain.faucet.GenesisState", GenesisState],
    ["/stwartchain.faucet.MsgUpdateTokensResponse", MsgUpdateTokensResponse],
    ["/stwartchain.faucet.MsgDeleteTokensResponse", MsgDeleteTokensResponse],
    ["/stwartchain.faucet.QueryAllTokensRequest", QueryAllTokensRequest],
    ["/stwartchain.faucet.QueryAllTokensResponse", QueryAllTokensResponse],
    ["/stwartchain.faucet.MsgUpdateTokens", MsgUpdateTokens],
    ["/stwartchain.faucet.MsgDeleteTokens", MsgDeleteTokens],
    ["/stwartchain.faucet.MsgIssueResponse", MsgIssueResponse],
    ["/stwartchain.faucet.QueryParamsResponse", QueryParamsResponse],
    ["/stwartchain.faucet.MsgUpdateParams", MsgUpdateParams],
    ["/stwartchain.faucet.MsgUpdateParamsResponse", MsgUpdateParamsResponse],
    ["/stwartchain.faucet.MsgCreateTokensResponse", MsgCreateTokensResponse],
    ["/stwartchain.faucet.Tokens", Tokens],
    ["/stwartchain.faucet.Params", Params],
    ["/stwartchain.faucet.MsgIssue", MsgIssue],
    ["/stwartchain.faucet.MsgCreateTokens", MsgCreateTokens],
    
];

export { msgTypes }
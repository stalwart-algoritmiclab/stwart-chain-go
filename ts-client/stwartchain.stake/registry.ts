import { GeneratedType } from "@cosmjs/proto-signing";
import { GenesisState } from "./types/stwartchain/stake/genesis";
import { QueryGetStakeRequest } from "./types/stwartchain/stake/query";
import { Stake } from "./types/stwartchain/stake/stake";
import { QueryGetStakeResponse } from "./types/stwartchain/stake/query";
import { QueryAllStakeRequest } from "./types/stwartchain/stake/query";
import { MsgUpdateParams } from "./types/stwartchain/stake/tx";
import { Params } from "./types/stwartchain/stake/params";
import { QueryParamsRequest } from "./types/stwartchain/stake/query";
import { QueryParamsResponse } from "./types/stwartchain/stake/query";
import { QueryAllStakeResponse } from "./types/stwartchain/stake/query";
import { MsgUpdateParamsResponse } from "./types/stwartchain/stake/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/stwartchain.stake.GenesisState", GenesisState],
    ["/stwartchain.stake.QueryGetStakeRequest", QueryGetStakeRequest],
    ["/stwartchain.stake.Stake", Stake],
    ["/stwartchain.stake.QueryGetStakeResponse", QueryGetStakeResponse],
    ["/stwartchain.stake.QueryAllStakeRequest", QueryAllStakeRequest],
    ["/stwartchain.stake.MsgUpdateParams", MsgUpdateParams],
    ["/stwartchain.stake.Params", Params],
    ["/stwartchain.stake.QueryParamsRequest", QueryParamsRequest],
    ["/stwartchain.stake.QueryParamsResponse", QueryParamsResponse],
    ["/stwartchain.stake.QueryAllStakeResponse", QueryAllStakeResponse],
    ["/stwartchain.stake.MsgUpdateParamsResponse", MsgUpdateParamsResponse],
    
];

export { msgTypes }
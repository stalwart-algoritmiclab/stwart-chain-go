import { GeneratedType } from "@cosmjs/proto-signing";
import { Params } from "./types/stwartchain/referral/params";
import { QueryParamsRequest } from "./types/stwartchain/referral/query";
import { QueryGetUserRequest } from "./types/stwartchain/referral/query";
import { MsgDeleteUser } from "./types/stwartchain/referral/tx";
import { MsgCreateUserResponse } from "./types/stwartchain/referral/tx";
import { GenesisState } from "./types/stwartchain/referral/genesis";
import { QueryAllUserRequest } from "./types/stwartchain/referral/query";
import { MsgUpdateParams } from "./types/stwartchain/referral/tx";
import { MsgDeleteUserResponse } from "./types/stwartchain/referral/tx";
import { User } from "./types/stwartchain/referral/user";
import { QueryParamsResponse } from "./types/stwartchain/referral/query";
import { QueryAllUserResponse } from "./types/stwartchain/referral/query";
import { MsgSetReferrer } from "./types/stwartchain/referral/tx";
import { MsgUpdateUser } from "./types/stwartchain/referral/tx";
import { MsgUpdateUserResponse } from "./types/stwartchain/referral/tx";
import { MsgSetReferrerResponse } from "./types/stwartchain/referral/tx";
import { MsgCreateUser } from "./types/stwartchain/referral/tx";
import { MsgUpdateParamsResponse } from "./types/stwartchain/referral/tx";
import { QueryGetUserResponse } from "./types/stwartchain/referral/query";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/stwartchain.referral.Params", Params],
    ["/stwartchain.referral.QueryParamsRequest", QueryParamsRequest],
    ["/stwartchain.referral.QueryGetUserRequest", QueryGetUserRequest],
    ["/stwartchain.referral.MsgDeleteUser", MsgDeleteUser],
    ["/stwartchain.referral.MsgCreateUserResponse", MsgCreateUserResponse],
    ["/stwartchain.referral.GenesisState", GenesisState],
    ["/stwartchain.referral.QueryAllUserRequest", QueryAllUserRequest],
    ["/stwartchain.referral.MsgUpdateParams", MsgUpdateParams],
    ["/stwartchain.referral.MsgDeleteUserResponse", MsgDeleteUserResponse],
    ["/stwartchain.referral.User", User],
    ["/stwartchain.referral.QueryParamsResponse", QueryParamsResponse],
    ["/stwartchain.referral.QueryAllUserResponse", QueryAllUserResponse],
    ["/stwartchain.referral.MsgSetReferrer", MsgSetReferrer],
    ["/stwartchain.referral.MsgUpdateUser", MsgUpdateUser],
    ["/stwartchain.referral.MsgUpdateUserResponse", MsgUpdateUserResponse],
    ["/stwartchain.referral.MsgSetReferrerResponse", MsgSetReferrerResponse],
    ["/stwartchain.referral.MsgCreateUser", MsgCreateUser],
    ["/stwartchain.referral.MsgUpdateParamsResponse", MsgUpdateParamsResponse],
    ["/stwartchain.referral.QueryGetUserResponse", QueryGetUserResponse],
    
];

export { msgTypes }
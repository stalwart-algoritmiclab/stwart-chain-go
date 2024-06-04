import { GeneratedType } from "@cosmjs/proto-signing";
import { QueryAllAddressesResponse } from "./types/stwartchain/feepolicy/query";
import { QueryGetTariffRequest } from "./types/stwartchain/feepolicy/query";
import { MsgUpdateParams } from "./types/stwartchain/feepolicy/tx";
import { MsgDeleteAddresses } from "./types/stwartchain/feepolicy/tx";
import { Address } from "./types/stwartchain/feepolicy/addresses";
import { QueryParamsResponse } from "./types/stwartchain/feepolicy/query";
import { QueryGetTariffResponse } from "./types/stwartchain/feepolicy/query";
import { QueryGetTariffsResponse } from "./types/stwartchain/feepolicy/query";
import { QueryAllTariffsResponse } from "./types/stwartchain/feepolicy/query";
import { MsgDeleteAddressesResponse } from "./types/stwartchain/feepolicy/tx";
import { MsgDeleteTariffs } from "./types/stwartchain/feepolicy/tx";
import { QueryAllTariffResponse } from "./types/stwartchain/feepolicy/query";
import { QueryGetAddressByIDRequest } from "./types/stwartchain/feepolicy/query";
import { QueryGetAddressResponse } from "./types/stwartchain/feepolicy/query";
import { MsgCreateTariffs } from "./types/stwartchain/feepolicy/tx";
import { MsgCreateTariffsResponse } from "./types/stwartchain/feepolicy/tx";
import { QueryGetAddressesResponse } from "./types/stwartchain/feepolicy/query";
import { QueryAllAddressesRequest } from "./types/stwartchain/feepolicy/query";
import { Tariff } from "./types/stwartchain/feepolicy/tariff";
import { Params } from "./types/stwartchain/feepolicy/params";
import { MsgUpdateAddresses } from "./types/stwartchain/feepolicy/tx";
import { MsgDeleteTariffsResponse } from "./types/stwartchain/feepolicy/tx";
import { QueryGetTariffsRequest } from "./types/stwartchain/feepolicy/query";
import { MsgUpdateParamsResponse } from "./types/stwartchain/feepolicy/tx";
import { MsgCreateAddressesResponse } from "./types/stwartchain/feepolicy/tx";
import { QueryAllTariffRequest } from "./types/stwartchain/feepolicy/query";
import { MsgUpdateTariffsResponse } from "./types/stwartchain/feepolicy/tx";
import { QueryGetAddressesRequest } from "./types/stwartchain/feepolicy/query";
import { MsgUpdateTariffs } from "./types/stwartchain/feepolicy/tx";
import { GenesisState } from "./types/stwartchain/feepolicy/genesis";
import { QueryAllTariffsRequest } from "./types/stwartchain/feepolicy/query";
import { Tariffs } from "./types/stwartchain/feepolicy/tariffs";
import { MsgCreateAddresses } from "./types/stwartchain/feepolicy/tx";
import { QueryParamsRequest } from "./types/stwartchain/feepolicy/query";
import { QueryGetAddressRequest } from "./types/stwartchain/feepolicy/query";
import { MsgUpdateAddressesResponse } from "./types/stwartchain/feepolicy/tx";
import { Fees } from "./types/stwartchain/feepolicy/fees";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/stwartchain.feepolicy.QueryAllAddressesResponse", QueryAllAddressesResponse],
    ["/stwartchain.feepolicy.QueryGetTariffRequest", QueryGetTariffRequest],
    ["/stwartchain.feepolicy.MsgUpdateParams", MsgUpdateParams],
    ["/stwartchain.feepolicy.MsgDeleteAddresses", MsgDeleteAddresses],
    ["/stwartchain.feepolicy.Address", Address],
    ["/stwartchain.feepolicy.QueryParamsResponse", QueryParamsResponse],
    ["/stwartchain.feepolicy.QueryGetTariffResponse", QueryGetTariffResponse],
    ["/stwartchain.feepolicy.QueryGetTariffsResponse", QueryGetTariffsResponse],
    ["/stwartchain.feepolicy.QueryAllTariffsResponse", QueryAllTariffsResponse],
    ["/stwartchain.feepolicy.MsgDeleteAddressesResponse", MsgDeleteAddressesResponse],
    ["/stwartchain.feepolicy.MsgDeleteTariffs", MsgDeleteTariffs],
    ["/stwartchain.feepolicy.QueryAllTariffResponse", QueryAllTariffResponse],
    ["/stwartchain.feepolicy.QueryGetAddressByIDRequest", QueryGetAddressByIDRequest],
    ["/stwartchain.feepolicy.QueryGetAddressResponse", QueryGetAddressResponse],
    ["/stwartchain.feepolicy.MsgCreateTariffs", MsgCreateTariffs],
    ["/stwartchain.feepolicy.MsgCreateTariffsResponse", MsgCreateTariffsResponse],
    ["/stwartchain.feepolicy.QueryGetAddressesResponse", QueryGetAddressesResponse],
    ["/stwartchain.feepolicy.QueryAllAddressesRequest", QueryAllAddressesRequest],
    ["/stwartchain.feepolicy.Tariff", Tariff],
    ["/stwartchain.feepolicy.Params", Params],
    ["/stwartchain.feepolicy.MsgUpdateAddresses", MsgUpdateAddresses],
    ["/stwartchain.feepolicy.MsgDeleteTariffsResponse", MsgDeleteTariffsResponse],
    ["/stwartchain.feepolicy.QueryGetTariffsRequest", QueryGetTariffsRequest],
    ["/stwartchain.feepolicy.MsgUpdateParamsResponse", MsgUpdateParamsResponse],
    ["/stwartchain.feepolicy.MsgCreateAddressesResponse", MsgCreateAddressesResponse],
    ["/stwartchain.feepolicy.QueryAllTariffRequest", QueryAllTariffRequest],
    ["/stwartchain.feepolicy.MsgUpdateTariffsResponse", MsgUpdateTariffsResponse],
    ["/stwartchain.feepolicy.QueryGetAddressesRequest", QueryGetAddressesRequest],
    ["/stwartchain.feepolicy.MsgUpdateTariffs", MsgUpdateTariffs],
    ["/stwartchain.feepolicy.GenesisState", GenesisState],
    ["/stwartchain.feepolicy.QueryAllTariffsRequest", QueryAllTariffsRequest],
    ["/stwartchain.feepolicy.Tariffs", Tariffs],
    ["/stwartchain.feepolicy.MsgCreateAddresses", MsgCreateAddresses],
    ["/stwartchain.feepolicy.QueryParamsRequest", QueryParamsRequest],
    ["/stwartchain.feepolicy.QueryGetAddressRequest", QueryGetAddressRequest],
    ["/stwartchain.feepolicy.MsgUpdateAddressesResponse", MsgUpdateAddressesResponse],
    ["/stwartchain.feepolicy.Fees", Fees],
    
];

export { msgTypes }
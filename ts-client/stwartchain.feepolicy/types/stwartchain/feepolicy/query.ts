/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../../cosmos/base/query/v1beta1/pagination";
import { Address } from "./addresses";
import { Params } from "./params";
import { Tariff } from "./tariff";
import { Tariffs } from "./tariffs";

export const protobufPackage = "stwartchain.feepolicy";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetAddressesRequest {
  id: number;
}

export interface QueryGetAddressesResponse {
  Addresses: Address | undefined;
}

export interface QueryAllAddressesRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllAddressesResponse {
  Addresses: Address[];
  pagination: PageResponse | undefined;
}

export interface QueryGetTariffRequest {
  denom: string;
}

export interface QueryGetTariffResponse {
  tariff: Tariff | undefined;
}

export interface QueryAllTariffRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllTariffResponse {
  tariff: Tariff[];
  pagination: PageResponse | undefined;
}

export interface QueryGetTariffsRequest {
  denom: string;
}

export interface QueryGetTariffsResponse {
  tariffs: Tariffs | undefined;
}

export interface QueryAllTariffsRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllTariffsResponse {
  tariffs: Tariffs[];
  pagination: PageResponse | undefined;
}

export interface QueryGetAddressByIDRequest {
  id: number;
}

export interface QueryGetAddressRequest {
  address: string;
}

export interface QueryGetAddressResponse {
  Address: Address | undefined;
}

function createBaseQueryParamsRequest(): QueryParamsRequest {
  return {};
}

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(_: any): QueryParamsRequest {
    return {};
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryParamsRequest>, I>>(base?: I): QueryParamsRequest {
    return QueryParamsRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryParamsRequest>, I>>(_: I): QueryParamsRequest {
    const message = createBaseQueryParamsRequest();
    return message;
  },
};

function createBaseQueryParamsResponse(): QueryParamsResponse {
  return { params: undefined };
}

export const QueryParamsResponse = {
  encode(message: QueryParamsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.params = Params.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    return { params: isSet(object.params) ? Params.fromJSON(object.params) : undefined };
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    if (message.params !== undefined) {
      obj.params = Params.toJSON(message.params);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryParamsResponse>, I>>(base?: I): QueryParamsResponse {
    return QueryParamsResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryParamsResponse>, I>>(object: I): QueryParamsResponse {
    const message = createBaseQueryParamsResponse();
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    return message;
  },
};

function createBaseQueryGetAddressesRequest(): QueryGetAddressesRequest {
  return { id: 0 };
}

export const QueryGetAddressesRequest = {
  encode(message: QueryGetAddressesRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetAddressesRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetAddressesRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.id = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): QueryGetAddressesRequest {
    return { id: isSet(object.id) ? Number(object.id) : 0 };
  },

  toJSON(message: QueryGetAddressesRequest): unknown {
    const obj: any = {};
    if (message.id !== 0) {
      obj.id = Math.round(message.id);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryGetAddressesRequest>, I>>(base?: I): QueryGetAddressesRequest {
    return QueryGetAddressesRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryGetAddressesRequest>, I>>(object: I): QueryGetAddressesRequest {
    const message = createBaseQueryGetAddressesRequest();
    message.id = object.id ?? 0;
    return message;
  },
};

function createBaseQueryGetAddressesResponse(): QueryGetAddressesResponse {
  return { Addresses: undefined };
}

export const QueryGetAddressesResponse = {
  encode(message: QueryGetAddressesResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.Addresses !== undefined) {
      Address.encode(message.Addresses, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetAddressesResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetAddressesResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.Addresses = Address.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): QueryGetAddressesResponse {
    return { Addresses: isSet(object.Addresses) ? Address.fromJSON(object.Addresses) : undefined };
  },

  toJSON(message: QueryGetAddressesResponse): unknown {
    const obj: any = {};
    if (message.Addresses !== undefined) {
      obj.Addresses = Address.toJSON(message.Addresses);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryGetAddressesResponse>, I>>(base?: I): QueryGetAddressesResponse {
    return QueryGetAddressesResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryGetAddressesResponse>, I>>(object: I): QueryGetAddressesResponse {
    const message = createBaseQueryGetAddressesResponse();
    message.Addresses = (object.Addresses !== undefined && object.Addresses !== null)
      ? Address.fromPartial(object.Addresses)
      : undefined;
    return message;
  },
};

function createBaseQueryAllAddressesRequest(): QueryAllAddressesRequest {
  return { pagination: undefined };
}

export const QueryAllAddressesRequest = {
  encode(message: QueryAllAddressesRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllAddressesRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllAddressesRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.pagination = PageRequest.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): QueryAllAddressesRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllAddressesRequest): unknown {
    const obj: any = {};
    if (message.pagination !== undefined) {
      obj.pagination = PageRequest.toJSON(message.pagination);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryAllAddressesRequest>, I>>(base?: I): QueryAllAddressesRequest {
    return QueryAllAddressesRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryAllAddressesRequest>, I>>(object: I): QueryAllAddressesRequest {
    const message = createBaseQueryAllAddressesRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllAddressesResponse(): QueryAllAddressesResponse {
  return { Addresses: [], pagination: undefined };
}

export const QueryAllAddressesResponse = {
  encode(message: QueryAllAddressesResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.Addresses) {
      Address.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllAddressesResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllAddressesResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.Addresses.push(Address.decode(reader, reader.uint32()));
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.pagination = PageResponse.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): QueryAllAddressesResponse {
    return {
      Addresses: Array.isArray(object?.Addresses) ? object.Addresses.map((e: any) => Address.fromJSON(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllAddressesResponse): unknown {
    const obj: any = {};
    if (message.Addresses?.length) {
      obj.Addresses = message.Addresses.map((e) => Address.toJSON(e));
    }
    if (message.pagination !== undefined) {
      obj.pagination = PageResponse.toJSON(message.pagination);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryAllAddressesResponse>, I>>(base?: I): QueryAllAddressesResponse {
    return QueryAllAddressesResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryAllAddressesResponse>, I>>(object: I): QueryAllAddressesResponse {
    const message = createBaseQueryAllAddressesResponse();
    message.Addresses = object.Addresses?.map((e) => Address.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetTariffRequest(): QueryGetTariffRequest {
  return { denom: "" };
}

export const QueryGetTariffRequest = {
  encode(message: QueryGetTariffRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.denom !== "") {
      writer.uint32(10).string(message.denom);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetTariffRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetTariffRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.denom = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): QueryGetTariffRequest {
    return { denom: isSet(object.denom) ? String(object.denom) : "" };
  },

  toJSON(message: QueryGetTariffRequest): unknown {
    const obj: any = {};
    if (message.denom !== "") {
      obj.denom = message.denom;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryGetTariffRequest>, I>>(base?: I): QueryGetTariffRequest {
    return QueryGetTariffRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryGetTariffRequest>, I>>(object: I): QueryGetTariffRequest {
    const message = createBaseQueryGetTariffRequest();
    message.denom = object.denom ?? "";
    return message;
  },
};

function createBaseQueryGetTariffResponse(): QueryGetTariffResponse {
  return { tariff: undefined };
}

export const QueryGetTariffResponse = {
  encode(message: QueryGetTariffResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.tariff !== undefined) {
      Tariff.encode(message.tariff, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetTariffResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetTariffResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.tariff = Tariff.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): QueryGetTariffResponse {
    return { tariff: isSet(object.tariff) ? Tariff.fromJSON(object.tariff) : undefined };
  },

  toJSON(message: QueryGetTariffResponse): unknown {
    const obj: any = {};
    if (message.tariff !== undefined) {
      obj.tariff = Tariff.toJSON(message.tariff);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryGetTariffResponse>, I>>(base?: I): QueryGetTariffResponse {
    return QueryGetTariffResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryGetTariffResponse>, I>>(object: I): QueryGetTariffResponse {
    const message = createBaseQueryGetTariffResponse();
    message.tariff = (object.tariff !== undefined && object.tariff !== null)
      ? Tariff.fromPartial(object.tariff)
      : undefined;
    return message;
  },
};

function createBaseQueryAllTariffRequest(): QueryAllTariffRequest {
  return { pagination: undefined };
}

export const QueryAllTariffRequest = {
  encode(message: QueryAllTariffRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllTariffRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllTariffRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.pagination = PageRequest.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): QueryAllTariffRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllTariffRequest): unknown {
    const obj: any = {};
    if (message.pagination !== undefined) {
      obj.pagination = PageRequest.toJSON(message.pagination);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryAllTariffRequest>, I>>(base?: I): QueryAllTariffRequest {
    return QueryAllTariffRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryAllTariffRequest>, I>>(object: I): QueryAllTariffRequest {
    const message = createBaseQueryAllTariffRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllTariffResponse(): QueryAllTariffResponse {
  return { tariff: [], pagination: undefined };
}

export const QueryAllTariffResponse = {
  encode(message: QueryAllTariffResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.tariff) {
      Tariff.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllTariffResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllTariffResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.tariff.push(Tariff.decode(reader, reader.uint32()));
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.pagination = PageResponse.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): QueryAllTariffResponse {
    return {
      tariff: Array.isArray(object?.tariff) ? object.tariff.map((e: any) => Tariff.fromJSON(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllTariffResponse): unknown {
    const obj: any = {};
    if (message.tariff?.length) {
      obj.tariff = message.tariff.map((e) => Tariff.toJSON(e));
    }
    if (message.pagination !== undefined) {
      obj.pagination = PageResponse.toJSON(message.pagination);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryAllTariffResponse>, I>>(base?: I): QueryAllTariffResponse {
    return QueryAllTariffResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryAllTariffResponse>, I>>(object: I): QueryAllTariffResponse {
    const message = createBaseQueryAllTariffResponse();
    message.tariff = object.tariff?.map((e) => Tariff.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetTariffsRequest(): QueryGetTariffsRequest {
  return { denom: "" };
}

export const QueryGetTariffsRequest = {
  encode(message: QueryGetTariffsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.denom !== "") {
      writer.uint32(10).string(message.denom);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetTariffsRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetTariffsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.denom = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): QueryGetTariffsRequest {
    return { denom: isSet(object.denom) ? String(object.denom) : "" };
  },

  toJSON(message: QueryGetTariffsRequest): unknown {
    const obj: any = {};
    if (message.denom !== "") {
      obj.denom = message.denom;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryGetTariffsRequest>, I>>(base?: I): QueryGetTariffsRequest {
    return QueryGetTariffsRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryGetTariffsRequest>, I>>(object: I): QueryGetTariffsRequest {
    const message = createBaseQueryGetTariffsRequest();
    message.denom = object.denom ?? "";
    return message;
  },
};

function createBaseQueryGetTariffsResponse(): QueryGetTariffsResponse {
  return { tariffs: undefined };
}

export const QueryGetTariffsResponse = {
  encode(message: QueryGetTariffsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.tariffs !== undefined) {
      Tariffs.encode(message.tariffs, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetTariffsResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetTariffsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.tariffs = Tariffs.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): QueryGetTariffsResponse {
    return { tariffs: isSet(object.tariffs) ? Tariffs.fromJSON(object.tariffs) : undefined };
  },

  toJSON(message: QueryGetTariffsResponse): unknown {
    const obj: any = {};
    if (message.tariffs !== undefined) {
      obj.tariffs = Tariffs.toJSON(message.tariffs);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryGetTariffsResponse>, I>>(base?: I): QueryGetTariffsResponse {
    return QueryGetTariffsResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryGetTariffsResponse>, I>>(object: I): QueryGetTariffsResponse {
    const message = createBaseQueryGetTariffsResponse();
    message.tariffs = (object.tariffs !== undefined && object.tariffs !== null)
      ? Tariffs.fromPartial(object.tariffs)
      : undefined;
    return message;
  },
};

function createBaseQueryAllTariffsRequest(): QueryAllTariffsRequest {
  return { pagination: undefined };
}

export const QueryAllTariffsRequest = {
  encode(message: QueryAllTariffsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllTariffsRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllTariffsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.pagination = PageRequest.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): QueryAllTariffsRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllTariffsRequest): unknown {
    const obj: any = {};
    if (message.pagination !== undefined) {
      obj.pagination = PageRequest.toJSON(message.pagination);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryAllTariffsRequest>, I>>(base?: I): QueryAllTariffsRequest {
    return QueryAllTariffsRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryAllTariffsRequest>, I>>(object: I): QueryAllTariffsRequest {
    const message = createBaseQueryAllTariffsRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllTariffsResponse(): QueryAllTariffsResponse {
  return { tariffs: [], pagination: undefined };
}

export const QueryAllTariffsResponse = {
  encode(message: QueryAllTariffsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.tariffs) {
      Tariffs.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllTariffsResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllTariffsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.tariffs.push(Tariffs.decode(reader, reader.uint32()));
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.pagination = PageResponse.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): QueryAllTariffsResponse {
    return {
      tariffs: Array.isArray(object?.tariffs) ? object.tariffs.map((e: any) => Tariffs.fromJSON(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllTariffsResponse): unknown {
    const obj: any = {};
    if (message.tariffs?.length) {
      obj.tariffs = message.tariffs.map((e) => Tariffs.toJSON(e));
    }
    if (message.pagination !== undefined) {
      obj.pagination = PageResponse.toJSON(message.pagination);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryAllTariffsResponse>, I>>(base?: I): QueryAllTariffsResponse {
    return QueryAllTariffsResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryAllTariffsResponse>, I>>(object: I): QueryAllTariffsResponse {
    const message = createBaseQueryAllTariffsResponse();
    message.tariffs = object.tariffs?.map((e) => Tariffs.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetAddressByIDRequest(): QueryGetAddressByIDRequest {
  return { id: 0 };
}

export const QueryGetAddressByIDRequest = {
  encode(message: QueryGetAddressByIDRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetAddressByIDRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetAddressByIDRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.id = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): QueryGetAddressByIDRequest {
    return { id: isSet(object.id) ? Number(object.id) : 0 };
  },

  toJSON(message: QueryGetAddressByIDRequest): unknown {
    const obj: any = {};
    if (message.id !== 0) {
      obj.id = Math.round(message.id);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryGetAddressByIDRequest>, I>>(base?: I): QueryGetAddressByIDRequest {
    return QueryGetAddressByIDRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryGetAddressByIDRequest>, I>>(object: I): QueryGetAddressByIDRequest {
    const message = createBaseQueryGetAddressByIDRequest();
    message.id = object.id ?? 0;
    return message;
  },
};

function createBaseQueryGetAddressRequest(): QueryGetAddressRequest {
  return { address: "" };
}

export const QueryGetAddressRequest = {
  encode(message: QueryGetAddressRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetAddressRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetAddressRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.address = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): QueryGetAddressRequest {
    return { address: isSet(object.address) ? String(object.address) : "" };
  },

  toJSON(message: QueryGetAddressRequest): unknown {
    const obj: any = {};
    if (message.address !== "") {
      obj.address = message.address;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryGetAddressRequest>, I>>(base?: I): QueryGetAddressRequest {
    return QueryGetAddressRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryGetAddressRequest>, I>>(object: I): QueryGetAddressRequest {
    const message = createBaseQueryGetAddressRequest();
    message.address = object.address ?? "";
    return message;
  },
};

function createBaseQueryGetAddressResponse(): QueryGetAddressResponse {
  return { Address: undefined };
}

export const QueryGetAddressResponse = {
  encode(message: QueryGetAddressResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.Address !== undefined) {
      Address.encode(message.Address, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetAddressResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetAddressResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.Address = Address.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): QueryGetAddressResponse {
    return { Address: isSet(object.Address) ? Address.fromJSON(object.Address) : undefined };
  },

  toJSON(message: QueryGetAddressResponse): unknown {
    const obj: any = {};
    if (message.Address !== undefined) {
      obj.Address = Address.toJSON(message.Address);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryGetAddressResponse>, I>>(base?: I): QueryGetAddressResponse {
    return QueryGetAddressResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryGetAddressResponse>, I>>(object: I): QueryGetAddressResponse {
    const message = createBaseQueryGetAddressResponse();
    message.Address = (object.Address !== undefined && object.Address !== null)
      ? Address.fromPartial(object.Address)
      : undefined;
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a list of Addresses items. */
  Addresses(request: QueryGetAddressesRequest): Promise<QueryGetAddressesResponse>;
  AddressesAll(request: QueryAllAddressesRequest): Promise<QueryAllAddressesResponse>;
  /** Queries a list of Tariff items. */
  Tariff(request: QueryGetTariffRequest): Promise<QueryGetTariffResponse>;
  TariffAll(request: QueryAllTariffRequest): Promise<QueryAllTariffResponse>;
  /** Queries a list of Tariffs items. */
  Tariffs(request: QueryGetTariffsRequest): Promise<QueryGetTariffsResponse>;
  TariffsAll(request: QueryAllTariffsRequest): Promise<QueryAllTariffsResponse>;
}

export const QueryServiceName = "stwartchain.feepolicy.Query";
export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || QueryServiceName;
    this.rpc = rpc;
    this.Params = this.Params.bind(this);
    this.Addresses = this.Addresses.bind(this);
    this.AddressesAll = this.AddressesAll.bind(this);
    this.Tariff = this.Tariff.bind(this);
    this.TariffAll = this.TariffAll.bind(this);
    this.Tariffs = this.Tariffs.bind(this);
    this.TariffsAll = this.TariffsAll.bind(this);
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(_m0.Reader.create(data)));
  }

  Addresses(request: QueryGetAddressesRequest): Promise<QueryGetAddressesResponse> {
    const data = QueryGetAddressesRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "Addresses", data);
    return promise.then((data) => QueryGetAddressesResponse.decode(_m0.Reader.create(data)));
  }

  AddressesAll(request: QueryAllAddressesRequest): Promise<QueryAllAddressesResponse> {
    const data = QueryAllAddressesRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "AddressesAll", data);
    return promise.then((data) => QueryAllAddressesResponse.decode(_m0.Reader.create(data)));
  }

  Tariff(request: QueryGetTariffRequest): Promise<QueryGetTariffResponse> {
    const data = QueryGetTariffRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "Tariff", data);
    return promise.then((data) => QueryGetTariffResponse.decode(_m0.Reader.create(data)));
  }

  TariffAll(request: QueryAllTariffRequest): Promise<QueryAllTariffResponse> {
    const data = QueryAllTariffRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "TariffAll", data);
    return promise.then((data) => QueryAllTariffResponse.decode(_m0.Reader.create(data)));
  }

  Tariffs(request: QueryGetTariffsRequest): Promise<QueryGetTariffsResponse> {
    const data = QueryGetTariffsRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "Tariffs", data);
    return promise.then((data) => QueryGetTariffsResponse.decode(_m0.Reader.create(data)));
  }

  TariffsAll(request: QueryAllTariffsRequest): Promise<QueryAllTariffsResponse> {
    const data = QueryAllTariffsRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "TariffsAll", data);
    return promise.then((data) => QueryAllTariffsResponse.decode(_m0.Reader.create(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

declare const self: any | undefined;
declare const window: any | undefined;
declare const global: any | undefined;
const tsProtoGlobalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new tsProtoGlobalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}

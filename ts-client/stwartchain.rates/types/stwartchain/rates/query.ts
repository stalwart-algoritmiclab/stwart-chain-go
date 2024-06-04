/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../../cosmos/base/query/v1beta1/pagination";
import { Addresses } from "./addresses";
import { Params } from "./params";
import { Rates } from "./rates";

export const protobufPackage = "stwartchain.rates";

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
  Addresses: Addresses | undefined;
}

export interface QueryAllAddressesRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllAddressesResponse {
  Addresses: Addresses[];
  pagination: PageResponse | undefined;
}

export interface QueryGetRatesRequest {
  denom: string;
}

export interface QueryGetRatesResponse {
  rates: Rates | undefined;
}

export interface QueryAllRatesRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllRatesResponse {
  rates: Rates[];
  pagination: PageResponse | undefined;
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
      Addresses.encode(message.Addresses, writer.uint32(10).fork()).ldelim();
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

          message.Addresses = Addresses.decode(reader, reader.uint32());
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
    return { Addresses: isSet(object.Addresses) ? Addresses.fromJSON(object.Addresses) : undefined };
  },

  toJSON(message: QueryGetAddressesResponse): unknown {
    const obj: any = {};
    if (message.Addresses !== undefined) {
      obj.Addresses = Addresses.toJSON(message.Addresses);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryGetAddressesResponse>, I>>(base?: I): QueryGetAddressesResponse {
    return QueryGetAddressesResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryGetAddressesResponse>, I>>(object: I): QueryGetAddressesResponse {
    const message = createBaseQueryGetAddressesResponse();
    message.Addresses = (object.Addresses !== undefined && object.Addresses !== null)
      ? Addresses.fromPartial(object.Addresses)
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
      Addresses.encode(v!, writer.uint32(10).fork()).ldelim();
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

          message.Addresses.push(Addresses.decode(reader, reader.uint32()));
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
      Addresses: Array.isArray(object?.Addresses) ? object.Addresses.map((e: any) => Addresses.fromJSON(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllAddressesResponse): unknown {
    const obj: any = {};
    if (message.Addresses?.length) {
      obj.Addresses = message.Addresses.map((e) => Addresses.toJSON(e));
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
    message.Addresses = object.Addresses?.map((e) => Addresses.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetRatesRequest(): QueryGetRatesRequest {
  return { denom: "" };
}

export const QueryGetRatesRequest = {
  encode(message: QueryGetRatesRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.denom !== "") {
      writer.uint32(10).string(message.denom);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetRatesRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetRatesRequest();
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

  fromJSON(object: any): QueryGetRatesRequest {
    return { denom: isSet(object.denom) ? String(object.denom) : "" };
  },

  toJSON(message: QueryGetRatesRequest): unknown {
    const obj: any = {};
    if (message.denom !== "") {
      obj.denom = message.denom;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryGetRatesRequest>, I>>(base?: I): QueryGetRatesRequest {
    return QueryGetRatesRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryGetRatesRequest>, I>>(object: I): QueryGetRatesRequest {
    const message = createBaseQueryGetRatesRequest();
    message.denom = object.denom ?? "";
    return message;
  },
};

function createBaseQueryGetRatesResponse(): QueryGetRatesResponse {
  return { rates: undefined };
}

export const QueryGetRatesResponse = {
  encode(message: QueryGetRatesResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.rates !== undefined) {
      Rates.encode(message.rates, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetRatesResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetRatesResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.rates = Rates.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): QueryGetRatesResponse {
    return { rates: isSet(object.rates) ? Rates.fromJSON(object.rates) : undefined };
  },

  toJSON(message: QueryGetRatesResponse): unknown {
    const obj: any = {};
    if (message.rates !== undefined) {
      obj.rates = Rates.toJSON(message.rates);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryGetRatesResponse>, I>>(base?: I): QueryGetRatesResponse {
    return QueryGetRatesResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryGetRatesResponse>, I>>(object: I): QueryGetRatesResponse {
    const message = createBaseQueryGetRatesResponse();
    message.rates = (object.rates !== undefined && object.rates !== null) ? Rates.fromPartial(object.rates) : undefined;
    return message;
  },
};

function createBaseQueryAllRatesRequest(): QueryAllRatesRequest {
  return { pagination: undefined };
}

export const QueryAllRatesRequest = {
  encode(message: QueryAllRatesRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllRatesRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllRatesRequest();
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

  fromJSON(object: any): QueryAllRatesRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllRatesRequest): unknown {
    const obj: any = {};
    if (message.pagination !== undefined) {
      obj.pagination = PageRequest.toJSON(message.pagination);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryAllRatesRequest>, I>>(base?: I): QueryAllRatesRequest {
    return QueryAllRatesRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryAllRatesRequest>, I>>(object: I): QueryAllRatesRequest {
    const message = createBaseQueryAllRatesRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllRatesResponse(): QueryAllRatesResponse {
  return { rates: [], pagination: undefined };
}

export const QueryAllRatesResponse = {
  encode(message: QueryAllRatesResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.rates) {
      Rates.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllRatesResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllRatesResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.rates.push(Rates.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllRatesResponse {
    return {
      rates: Array.isArray(object?.rates) ? object.rates.map((e: any) => Rates.fromJSON(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllRatesResponse): unknown {
    const obj: any = {};
    if (message.rates?.length) {
      obj.rates = message.rates.map((e) => Rates.toJSON(e));
    }
    if (message.pagination !== undefined) {
      obj.pagination = PageResponse.toJSON(message.pagination);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryAllRatesResponse>, I>>(base?: I): QueryAllRatesResponse {
    return QueryAllRatesResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryAllRatesResponse>, I>>(object: I): QueryAllRatesResponse {
    const message = createBaseQueryAllRatesResponse();
    message.rates = object.rates?.map((e) => Rates.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
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
  /** Queries a list of Rates items. */
  Rates(request: QueryGetRatesRequest): Promise<QueryGetRatesResponse>;
  RatesAll(request: QueryAllRatesRequest): Promise<QueryAllRatesResponse>;
}

export const QueryServiceName = "stwartchain.rates.Query";
export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || QueryServiceName;
    this.rpc = rpc;
    this.Params = this.Params.bind(this);
    this.Addresses = this.Addresses.bind(this);
    this.AddressesAll = this.AddressesAll.bind(this);
    this.Rates = this.Rates.bind(this);
    this.RatesAll = this.RatesAll.bind(this);
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

  Rates(request: QueryGetRatesRequest): Promise<QueryGetRatesResponse> {
    const data = QueryGetRatesRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "Rates", data);
    return promise.then((data) => QueryGetRatesResponse.decode(_m0.Reader.create(data)));
  }

  RatesAll(request: QueryAllRatesRequest): Promise<QueryAllRatesResponse> {
    const data = QueryAllRatesRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "RatesAll", data);
    return promise.then((data) => QueryAllRatesResponse.decode(_m0.Reader.create(data)));
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

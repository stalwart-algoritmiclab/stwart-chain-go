/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Params } from "./params";

export const protobufPackage = "stwartchain.referral";

/** MsgUpdateParams is the Msg/UpdateParams request type. */
export interface MsgUpdateParams {
  /** authority is the address that controls the module (defaults to x/gov unless overwritten). */
  authority: string;
  /** NOTE: All parameters must be supplied. */
  params: Params | undefined;
}

/**
 * MsgUpdateParamsResponse defines the response structure for executing a
 * MsgUpdateParams message.
 */
export interface MsgUpdateParamsResponse {
}

export interface MsgCreateUser {
  creator: string;
  accountAddress: string;
  referrer: string;
  referrals: string[];
}

export interface MsgCreateUserResponse {
}

export interface MsgUpdateUser {
  creator: string;
  accountAddress: string;
  referrer: string;
  referrals: string[];
}

export interface MsgUpdateUserResponse {
}

export interface MsgDeleteUser {
  creator: string;
  accountAddress: string;
}

export interface MsgDeleteUserResponse {
}

export interface MsgSetReferrer {
  creator: string;
  referrerAddress: string;
  referralAddress: string;
}

export interface MsgSetReferrerResponse {
}

function createBaseMsgUpdateParams(): MsgUpdateParams {
  return { authority: "", params: undefined };
}

export const MsgUpdateParams = {
  encode(message: MsgUpdateParams, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.authority !== "") {
      writer.uint32(10).string(message.authority);
    }
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateParams {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateParams();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.authority = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
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

  fromJSON(object: any): MsgUpdateParams {
    return {
      authority: isSet(object.authority) ? String(object.authority) : "",
      params: isSet(object.params) ? Params.fromJSON(object.params) : undefined,
    };
  },

  toJSON(message: MsgUpdateParams): unknown {
    const obj: any = {};
    if (message.authority !== "") {
      obj.authority = message.authority;
    }
    if (message.params !== undefined) {
      obj.params = Params.toJSON(message.params);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<MsgUpdateParams>, I>>(base?: I): MsgUpdateParams {
    return MsgUpdateParams.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<MsgUpdateParams>, I>>(object: I): MsgUpdateParams {
    const message = createBaseMsgUpdateParams();
    message.authority = object.authority ?? "";
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    return message;
  },
};

function createBaseMsgUpdateParamsResponse(): MsgUpdateParamsResponse {
  return {};
}

export const MsgUpdateParamsResponse = {
  encode(_: MsgUpdateParamsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateParamsResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateParamsResponse();
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

  fromJSON(_: any): MsgUpdateParamsResponse {
    return {};
  },

  toJSON(_: MsgUpdateParamsResponse): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<MsgUpdateParamsResponse>, I>>(base?: I): MsgUpdateParamsResponse {
    return MsgUpdateParamsResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<MsgUpdateParamsResponse>, I>>(_: I): MsgUpdateParamsResponse {
    const message = createBaseMsgUpdateParamsResponse();
    return message;
  },
};

function createBaseMsgCreateUser(): MsgCreateUser {
  return { creator: "", accountAddress: "", referrer: "", referrals: [] };
}

export const MsgCreateUser = {
  encode(message: MsgCreateUser, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.accountAddress !== "") {
      writer.uint32(18).string(message.accountAddress);
    }
    if (message.referrer !== "") {
      writer.uint32(26).string(message.referrer);
    }
    for (const v of message.referrals) {
      writer.uint32(34).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateUser {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateUser();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.creator = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.accountAddress = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.referrer = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.referrals.push(reader.string());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): MsgCreateUser {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      accountAddress: isSet(object.accountAddress) ? String(object.accountAddress) : "",
      referrer: isSet(object.referrer) ? String(object.referrer) : "",
      referrals: Array.isArray(object?.referrals) ? object.referrals.map((e: any) => String(e)) : [],
    };
  },

  toJSON(message: MsgCreateUser): unknown {
    const obj: any = {};
    if (message.creator !== "") {
      obj.creator = message.creator;
    }
    if (message.accountAddress !== "") {
      obj.accountAddress = message.accountAddress;
    }
    if (message.referrer !== "") {
      obj.referrer = message.referrer;
    }
    if (message.referrals?.length) {
      obj.referrals = message.referrals;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<MsgCreateUser>, I>>(base?: I): MsgCreateUser {
    return MsgCreateUser.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<MsgCreateUser>, I>>(object: I): MsgCreateUser {
    const message = createBaseMsgCreateUser();
    message.creator = object.creator ?? "";
    message.accountAddress = object.accountAddress ?? "";
    message.referrer = object.referrer ?? "";
    message.referrals = object.referrals?.map((e) => e) || [];
    return message;
  },
};

function createBaseMsgCreateUserResponse(): MsgCreateUserResponse {
  return {};
}

export const MsgCreateUserResponse = {
  encode(_: MsgCreateUserResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateUserResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateUserResponse();
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

  fromJSON(_: any): MsgCreateUserResponse {
    return {};
  },

  toJSON(_: MsgCreateUserResponse): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<MsgCreateUserResponse>, I>>(base?: I): MsgCreateUserResponse {
    return MsgCreateUserResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<MsgCreateUserResponse>, I>>(_: I): MsgCreateUserResponse {
    const message = createBaseMsgCreateUserResponse();
    return message;
  },
};

function createBaseMsgUpdateUser(): MsgUpdateUser {
  return { creator: "", accountAddress: "", referrer: "", referrals: [] };
}

export const MsgUpdateUser = {
  encode(message: MsgUpdateUser, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.accountAddress !== "") {
      writer.uint32(18).string(message.accountAddress);
    }
    if (message.referrer !== "") {
      writer.uint32(26).string(message.referrer);
    }
    for (const v of message.referrals) {
      writer.uint32(34).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateUser {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateUser();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.creator = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.accountAddress = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.referrer = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.referrals.push(reader.string());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateUser {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      accountAddress: isSet(object.accountAddress) ? String(object.accountAddress) : "",
      referrer: isSet(object.referrer) ? String(object.referrer) : "",
      referrals: Array.isArray(object?.referrals) ? object.referrals.map((e: any) => String(e)) : [],
    };
  },

  toJSON(message: MsgUpdateUser): unknown {
    const obj: any = {};
    if (message.creator !== "") {
      obj.creator = message.creator;
    }
    if (message.accountAddress !== "") {
      obj.accountAddress = message.accountAddress;
    }
    if (message.referrer !== "") {
      obj.referrer = message.referrer;
    }
    if (message.referrals?.length) {
      obj.referrals = message.referrals;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<MsgUpdateUser>, I>>(base?: I): MsgUpdateUser {
    return MsgUpdateUser.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<MsgUpdateUser>, I>>(object: I): MsgUpdateUser {
    const message = createBaseMsgUpdateUser();
    message.creator = object.creator ?? "";
    message.accountAddress = object.accountAddress ?? "";
    message.referrer = object.referrer ?? "";
    message.referrals = object.referrals?.map((e) => e) || [];
    return message;
  },
};

function createBaseMsgUpdateUserResponse(): MsgUpdateUserResponse {
  return {};
}

export const MsgUpdateUserResponse = {
  encode(_: MsgUpdateUserResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateUserResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateUserResponse();
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

  fromJSON(_: any): MsgUpdateUserResponse {
    return {};
  },

  toJSON(_: MsgUpdateUserResponse): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<MsgUpdateUserResponse>, I>>(base?: I): MsgUpdateUserResponse {
    return MsgUpdateUserResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<MsgUpdateUserResponse>, I>>(_: I): MsgUpdateUserResponse {
    const message = createBaseMsgUpdateUserResponse();
    return message;
  },
};

function createBaseMsgDeleteUser(): MsgDeleteUser {
  return { creator: "", accountAddress: "" };
}

export const MsgDeleteUser = {
  encode(message: MsgDeleteUser, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.accountAddress !== "") {
      writer.uint32(18).string(message.accountAddress);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDeleteUser {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeleteUser();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.creator = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.accountAddress = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteUser {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      accountAddress: isSet(object.accountAddress) ? String(object.accountAddress) : "",
    };
  },

  toJSON(message: MsgDeleteUser): unknown {
    const obj: any = {};
    if (message.creator !== "") {
      obj.creator = message.creator;
    }
    if (message.accountAddress !== "") {
      obj.accountAddress = message.accountAddress;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<MsgDeleteUser>, I>>(base?: I): MsgDeleteUser {
    return MsgDeleteUser.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<MsgDeleteUser>, I>>(object: I): MsgDeleteUser {
    const message = createBaseMsgDeleteUser();
    message.creator = object.creator ?? "";
    message.accountAddress = object.accountAddress ?? "";
    return message;
  },
};

function createBaseMsgDeleteUserResponse(): MsgDeleteUserResponse {
  return {};
}

export const MsgDeleteUserResponse = {
  encode(_: MsgDeleteUserResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDeleteUserResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeleteUserResponse();
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

  fromJSON(_: any): MsgDeleteUserResponse {
    return {};
  },

  toJSON(_: MsgDeleteUserResponse): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<MsgDeleteUserResponse>, I>>(base?: I): MsgDeleteUserResponse {
    return MsgDeleteUserResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<MsgDeleteUserResponse>, I>>(_: I): MsgDeleteUserResponse {
    const message = createBaseMsgDeleteUserResponse();
    return message;
  },
};

function createBaseMsgSetReferrer(): MsgSetReferrer {
  return { creator: "", referrerAddress: "", referralAddress: "" };
}

export const MsgSetReferrer = {
  encode(message: MsgSetReferrer, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.referrerAddress !== "") {
      writer.uint32(18).string(message.referrerAddress);
    }
    if (message.referralAddress !== "") {
      writer.uint32(26).string(message.referralAddress);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSetReferrer {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSetReferrer();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.creator = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.referrerAddress = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.referralAddress = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): MsgSetReferrer {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      referrerAddress: isSet(object.referrerAddress) ? String(object.referrerAddress) : "",
      referralAddress: isSet(object.referralAddress) ? String(object.referralAddress) : "",
    };
  },

  toJSON(message: MsgSetReferrer): unknown {
    const obj: any = {};
    if (message.creator !== "") {
      obj.creator = message.creator;
    }
    if (message.referrerAddress !== "") {
      obj.referrerAddress = message.referrerAddress;
    }
    if (message.referralAddress !== "") {
      obj.referralAddress = message.referralAddress;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<MsgSetReferrer>, I>>(base?: I): MsgSetReferrer {
    return MsgSetReferrer.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<MsgSetReferrer>, I>>(object: I): MsgSetReferrer {
    const message = createBaseMsgSetReferrer();
    message.creator = object.creator ?? "";
    message.referrerAddress = object.referrerAddress ?? "";
    message.referralAddress = object.referralAddress ?? "";
    return message;
  },
};

function createBaseMsgSetReferrerResponse(): MsgSetReferrerResponse {
  return {};
}

export const MsgSetReferrerResponse = {
  encode(_: MsgSetReferrerResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSetReferrerResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSetReferrerResponse();
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

  fromJSON(_: any): MsgSetReferrerResponse {
    return {};
  },

  toJSON(_: MsgSetReferrerResponse): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<MsgSetReferrerResponse>, I>>(base?: I): MsgSetReferrerResponse {
    return MsgSetReferrerResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<MsgSetReferrerResponse>, I>>(_: I): MsgSetReferrerResponse {
    const message = createBaseMsgSetReferrerResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  /**
   * UpdateParams defines a (governance) operation for updating the module
   * parameters. The authority defaults to the x/gov module account.
   */
  UpdateParams(request: MsgUpdateParams): Promise<MsgUpdateParamsResponse>;
  CreateUser(request: MsgCreateUser): Promise<MsgCreateUserResponse>;
  UpdateUser(request: MsgUpdateUser): Promise<MsgUpdateUserResponse>;
  DeleteUser(request: MsgDeleteUser): Promise<MsgDeleteUserResponse>;
  SetReferrer(request: MsgSetReferrer): Promise<MsgSetReferrerResponse>;
}

export const MsgServiceName = "stwartchain.referral.Msg";
export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || MsgServiceName;
    this.rpc = rpc;
    this.UpdateParams = this.UpdateParams.bind(this);
    this.CreateUser = this.CreateUser.bind(this);
    this.UpdateUser = this.UpdateUser.bind(this);
    this.DeleteUser = this.DeleteUser.bind(this);
    this.SetReferrer = this.SetReferrer.bind(this);
  }
  UpdateParams(request: MsgUpdateParams): Promise<MsgUpdateParamsResponse> {
    const data = MsgUpdateParams.encode(request).finish();
    const promise = this.rpc.request(this.service, "UpdateParams", data);
    return promise.then((data) => MsgUpdateParamsResponse.decode(_m0.Reader.create(data)));
  }

  CreateUser(request: MsgCreateUser): Promise<MsgCreateUserResponse> {
    const data = MsgCreateUser.encode(request).finish();
    const promise = this.rpc.request(this.service, "CreateUser", data);
    return promise.then((data) => MsgCreateUserResponse.decode(_m0.Reader.create(data)));
  }

  UpdateUser(request: MsgUpdateUser): Promise<MsgUpdateUserResponse> {
    const data = MsgUpdateUser.encode(request).finish();
    const promise = this.rpc.request(this.service, "UpdateUser", data);
    return promise.then((data) => MsgUpdateUserResponse.decode(_m0.Reader.create(data)));
  }

  DeleteUser(request: MsgDeleteUser): Promise<MsgDeleteUserResponse> {
    const data = MsgDeleteUser.encode(request).finish();
    const promise = this.rpc.request(this.service, "DeleteUser", data);
    return promise.then((data) => MsgDeleteUserResponse.decode(_m0.Reader.create(data)));
  }

  SetReferrer(request: MsgSetReferrer): Promise<MsgSetReferrerResponse> {
    const data = MsgSetReferrer.encode(request).finish();
    const promise = this.rpc.request(this.service, "SetReferrer", data);
    return promise.then((data) => MsgSetReferrerResponse.decode(_m0.Reader.create(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}

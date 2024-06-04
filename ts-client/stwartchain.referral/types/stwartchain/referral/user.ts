/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "stwartchain.referral";

export interface User {
  accountAddress: string;
  referrer: string;
  referrals: string[];
}

function createBaseUser(): User {
  return { accountAddress: "", referrer: "", referrals: [] };
}

export const User = {
  encode(message: User, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.accountAddress !== "") {
      writer.uint32(10).string(message.accountAddress);
    }
    if (message.referrer !== "") {
      writer.uint32(18).string(message.referrer);
    }
    for (const v of message.referrals) {
      writer.uint32(26).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): User {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUser();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.accountAddress = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.referrer = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
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

  fromJSON(object: any): User {
    return {
      accountAddress: isSet(object.accountAddress) ? String(object.accountAddress) : "",
      referrer: isSet(object.referrer) ? String(object.referrer) : "",
      referrals: Array.isArray(object?.referrals) ? object.referrals.map((e: any) => String(e)) : [],
    };
  },

  toJSON(message: User): unknown {
    const obj: any = {};
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

  create<I extends Exact<DeepPartial<User>, I>>(base?: I): User {
    return User.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<User>, I>>(object: I): User {
    const message = createBaseUser();
    message.accountAddress = object.accountAddress ?? "";
    message.referrer = object.referrer ?? "";
    message.referrals = object.referrals?.map((e) => e) || [];
    return message;
  },
};

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

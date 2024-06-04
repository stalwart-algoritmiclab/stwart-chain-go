/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { DailyStats } from "./daily_stats";

export const protobufPackage = "stwartchain.core";

export interface Stats {
  date: string;
  dailyStats: DailyStats | undefined;
}

function createBaseStats(): Stats {
  return { date: "", dailyStats: undefined };
}

export const Stats = {
  encode(message: Stats, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.date !== "") {
      writer.uint32(10).string(message.date);
    }
    if (message.dailyStats !== undefined) {
      DailyStats.encode(message.dailyStats, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Stats {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseStats();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.date = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.dailyStats = DailyStats.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Stats {
    return {
      date: isSet(object.date) ? String(object.date) : "",
      dailyStats: isSet(object.dailyStats) ? DailyStats.fromJSON(object.dailyStats) : undefined,
    };
  },

  toJSON(message: Stats): unknown {
    const obj: any = {};
    if (message.date !== "") {
      obj.date = message.date;
    }
    if (message.dailyStats !== undefined) {
      obj.dailyStats = DailyStats.toJSON(message.dailyStats);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<Stats>, I>>(base?: I): Stats {
    return Stats.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<Stats>, I>>(object: I): Stats {
    const message = createBaseStats();
    message.date = object.date ?? "";
    message.dailyStats = (object.dailyStats !== undefined && object.dailyStats !== null)
      ? DailyStats.fromPartial(object.dailyStats)
      : undefined;
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

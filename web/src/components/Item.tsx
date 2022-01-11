import React from 'react';

type RecordContent = {
  String: string;
  Valid: boolean;
}

export type SensorRecord = {
  Id: number;
  Date: Date;
  Beacon: string;
  Name: string;
  RecordType: number;
  Record: RecordContent;
};

type Props = {
  records: SensorRecord[];
} & typeof defaultProps;

const defaultProps = {};

export const Item = ({ records }: Props): JSX.Element => {

  return (
    <div
      onClick={() => { console.log(records) }}
      style={{
        display: 'flex',
        flexDirection: 'column',
        justifyContent: 'space-between',
        alignItems: 'center',
        border: 'thin solid white',
        overflow: 'auto',
        cursor: 'pointer',
      }}
    >
      {/* TODO: Handle overflowing */}
      {/* TODO: Update props and display accordingly */}
      {/* Item Header */}
      <div
        style={{
          display: 'flex',
          flexDirection: 'column',
          justifyContent: 'center',
          alignItems: 'center',
        }}
      >
        <div
          style={{
            fontSize: 20,
          }}
        >
          { records[0].Name }
        </div>
        <div
          style={{
            fontSize: 13,
          }}
        >
          { records[0].Beacon }
        </div>
      </div>
      {/* Item Content */}
      <div
        style={{
          fontSize: 44,
          fontWeight: 'bold',
        }}
      >
        { records[0].Record.String }
      </div>
      {/* Item Footer */}
      <div
        style={{

        }}
      />
    </div>
  );
}
Item.defaultProps = defaultProps;

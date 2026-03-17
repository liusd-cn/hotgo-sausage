import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';

export class State {
  public id = 0; // id
  public name = ''; // name
  public type = 0; // type
  public saleTime = ''; // sale_time

  constructor(state?: Partial<State>) {
    if (state) {
      Object.assign(this, state);
    }
  }
}

export function newState(state: State | Record<string, any> | null): State {
  if (state !== null) {
    if (state instanceof State) {
      return cloneDeep(state);
    }
    return new State(state);
  }
  return new State();
}

// 表单验证规则

// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  {
    field: 'id',
    component: 'NInputNumber',
    label: 'id',
    componentProps: {
      placeholder: '请输入id',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
]);

// 表格列
export const columns = [
  {
    title: 'id',
    key: 'id',
    align: 'left',
    width: -1,
  },
  {
    title: 'name',
    key: 'name',
    align: 'left',
    width: -1,
  },
  {
    title: 'type',
    key: 'type',
    align: 'left',
    width: -1,
  },
  {
    title: 'sale_time',
    key: 'saleTime',
    align: 'left',
    width: -1,
  },
];
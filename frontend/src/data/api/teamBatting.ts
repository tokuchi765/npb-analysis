import { rest } from '../rest';
import { MaxTeamBattingResponse, MinTeamBattingResponse } from '../type';

const getMaxTeamBatting = async (): Promise<MaxTeamBattingResponse> => {
  try {
    const { data } = await rest.get<MaxTeamBattingResponse>('/team/batting/max');
    return data;
  } catch (error: any) {
    throw new Error(error);
  }
};

const getMinTeamBatting = async (): Promise<MinTeamBattingResponse> => {
  try {
    const { data } = await rest.get<MinTeamBattingResponse>('/team/batting/min');
    return data;
  } catch (error: any) {
    throw new Error(error);
  }
};

export { getMaxTeamBatting, getMinTeamBatting };

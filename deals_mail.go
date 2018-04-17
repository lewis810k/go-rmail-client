package rmail

func (client *Client) CreateDeal(dealId int) {
	query := `
		mutation sendCreateDealMail($id: Int!) {
			send_mail {
				deals {
					create_deal(deal_id: $id)
				}
			}
		}
	`
	vars := map[string]interface{}{
		"id": dealId,
	}
	client.sendMail(query, vars)
}

func (client *Client) AddUserToDeal(dealId int, userId int) {
	query := `
		mutation sendAddUserToDealMail($deal_id: Int!, $user_id: Int!) {
			send_mail {
				deals {
					add_user_to_deal(deal_id: $deal_id, user_id: $user_id)
				}
			}
		}
	`
	vars := map[string]interface{}{
		"deal_id": dealId,
		"user_id": userId,
	}
	client.sendMail(query, vars)
}

func (client *Client) RemoveUserFromDeal(dealId int, userId int) {
	query := `
		mutation sendRemoveUserFromDealMail($deal_id: Int!, $user_id: Int!) {
			send_mail {
				deals {
					remove_user_from_deal(deal_id: $deal_id, user_id: $user_id)
				}
			}
		}
	`
	vars := map[string]interface{}{
		"deal_id": dealId,
		"user_id": userId,
	}
	client.sendMail(query, vars)
}
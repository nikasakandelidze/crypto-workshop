# goofy coin main idea:
	1)only goofy can produce new coins
	2) every user has public keys associated with them
	3) every user can transfer coin to another user by specifying a statement in datastructure
	payload and signing it with public key.
## Generating and owning coins
	It's a centralized crypto currency. There is one centralized authority called goofy
	who can generate new coins. Upon generating new coin goofy signs it with his public
	key, which means that now he owns this coin.
	Coins data structure:  COIN_1
		[ signed by goofy's pulbic key ] signature
		[     CreateCoin(UniqueId)     ] payload

## Transactions of the coins
	Goofy and any coin owner can transfer their coins to other participants.
	mechanism is next:
		1) create new structure which is signed by coin owner and in payload
		has data about to whom to transact this coin. like next:
				COIN_2
			[ signed by goofy's public key ] signature]
			[ Pay to public key Alice. Hash pointer to COIN_1]
		Alice can prove that she owns a coin by presenting strucutre above, since 
		it was signed to goofy and previosuly owned by him, anyone can validate this ds with goofy.
## problem
	1) Central authroity of coin generation. If goofy would like to do something malicous or bad no one would go agains him.
	2) Another problem with this coin is that whole tranmsaction log/datastructure is not syncrhonized between
	parties  and every current coin user has only his/her copuy of log. Which means that at some point
	coin owner can create 2 structures sign both of them and in the payload specify for each of them
	different target public figure. Both of them will recieve this structure and both of them will assume that
	they now own the coin but in reality original cown owner had only one coin. THis is called double spending.	 

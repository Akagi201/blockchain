/**
 *  @file
 *  @copyright defined in eos/LICENSE.txt
 */

#include <eosiolib/eosio.hpp>
#include <string>

namespace eosio {
   using ipfshash_t = std::string;

   class bmail : public contract {
      public:
         bmail( account_name self ):contract(self) {}
         
	 //@abi action
         void sendmail( account_name  sender, 
                        const vector<account_name>& receivers,
                        const vector<ipfshash_t>& mailhashs );

      private:
	 //@abi table dbmail i64
         struct mail {
 	    mail() { mailhash.resize(34); }
            uint64_t      id;
            account_name  sender;               
            account_name  receiver;
            ipfshash_t    mailhash;
	    uint32_t      status = 0;            
            time          create_date;
	    time          update_date;
		 
	    uint64_t primary_key()const { return id; }
		 
            EOSLIB_SERIALIZE(mail, (id)(sender)(receiver)(mailhash)(status)(create_date)(update_date) )
	};

        typedef eosio::multi_index<N(dbmail), mail> mail_index;
   };

} /// namespace eosio

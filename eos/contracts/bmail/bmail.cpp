/**
 *  @file
 *  @copyright defined in eos/LICENSE.txt
 */

#include <bmail.hpp>

using namespace eosio;

void bmail::sendmail( account_name sender, 
                      const vector<account_name>& receivers,
                      const vector<ipfshash_t>& mailhashs )
{  
    require_auth( sender );
    
    eosio_assert(receivers.size() == mailhashs.size(), "nº receivers must match nº of mailhashs");

    for(int i = 0; i < receivers.size(); i++) {
       auto receiver = receivers[i];
       auto mailhash = mailhashs[i];
      
       eosio_assert(mailhash.size() == 34, "mailhash needs to be a ipfs hash with 34 length");

       mail_index sender_mails(_self, sender);
       sender_mails.emplace( sender, [&]( auto& a ) {
          a.id = sender_mails.available_primary_key();
          a.sender = sender;
          a.receiver = receiver;
          a.mailhash = mailhash;
          a.create_date = a.update_date = now();
       });
       
       mail_index receiver_mails(_self, receiver);
       receiver_mails.emplace( sender, [&]( auto& a ) {
         a.id = receiver_mails.available_primary_key();
         a.sender = sender;
         a.receiver = receiver;
         a.mailhash = mailhash;
         a.create_date = a.update_date = now();
      });
    }
}

EOSIO_ABI( bmail, (sendmail) )
